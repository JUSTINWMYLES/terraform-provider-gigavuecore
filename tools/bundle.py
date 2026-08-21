#!/usr/bin/env python3
"""Bundle external JSON Schema refs into a single OpenAPI spec with local refs."""
import os, json, re, glob
from collections import defaultdict
import yaml

ROOT = os.path.join(os.path.dirname(os.path.dirname(os.path.abspath(__file__))), 'spec')
MAIN = os.path.join(ROOT, 'openapi.fm.yaml')

# ---------- Load all files ----------
files = {}
for f in glob.glob(f'{ROOT}/source/**/*.json', recursive=True):
    try:
        files[os.path.abspath(f)] = json.load(open(f))
    except Exception as e:
        print(f"SKIP {f}: {e}")

def canon(p):
    return os.path.abspath(os.path.normpath(p))

# ---------- Flat def map per file ----------
flat_defs = {}

def collect_defs(obj, file_abs, out):
    """Collect definitions: top-level + nested 'definitions' blocks + id-anchored sub-schemas."""
    if isinstance(obj, dict):
        if 'definitions' in obj and isinstance(obj['definitions'], dict):
            for n, s in obj['definitions'].items():
                if n not in out:
                    out[n] = s
                collect_defs(s, file_abs, out)
        # id-anchored sub-schemas (draft-04): {"id": "#/definitions/Name", ...}
        for n, v in obj.items():
            if isinstance(v, dict) and n not in ('definitions',):
                vid = v.get('id')
                if isinstance(vid, str) and vid.startswith('#/definitions/'):
                    dn = vid[len('#/definitions/'):]
                    if dn not in out:
                        out[dn] = v
                collect_defs(v, file_abs, out)
    elif isinstance(obj, list):
        for v in obj:
            collect_defs(v, file_abs, out)

for f, d in files.items():
    fd = {}
    collect_defs(d, f, fd)
    flat_defs[f] = fd

# ---------- Collision-safe name mapping ----------
name_owners = defaultdict(list)
for f, fd in flat_defs.items():
    for n in fd:
        name_owners[n].append(f)
collisions = {n: vs for n, vs in name_owners.items() if len(vs) > 1}

def path_key(f):
    rel = os.path.relpath(f, ROOT)
    parts = rel.split('/')
    return '_'.join(parts[:-1] + [parts[-1][:-5]])

def map_name(f, name):
    if name not in collisions:
        return name
    return f"{path_key(f)}_{name}"

# ---------- Upstream broken-ref fixes ----------
FIXES = {
    ('source/schemas-fm/licensing.json', 'definitions/ListOfIntegersReturn'):
        {'type': 'array', 'items': {'type': 'integer'}},
    ('source/schemas-fm/licensing.json', 'definitions/ListOfStringsReturn'):
        {'type': 'array', 'items': {'type': 'string'}},
    ('source/schemas-fm/flows.json', 'definitions/AppConfigRulesAppConfigRulesSingleQueryResponse'):
        ('POINTER', 'source/schemas-fm/flows.json', 'definitions/AppConfigRulesSingleQueryResponse'),
    ('source/schemas-fm/sourcesAndRules.json', 'definitions/GigaAfmRules'):
        ('POINTER', 'source/schemas-fm/sourcesAndRules.json', 'definitions/sourcesAndRules'),
}
FIX_KEYS = {}
for (fp, ptr), val in FIXES.items():
    key = (canon(os.path.join(ROOT, fp)), ptr)
    FIX_KEYS[key] = val

# ---------- Component registry ----------
components = {}
pending = []

def get_pointer_schema(file_abs, pointer):
    if not pointer.startswith('definitions/'):
        return None
    rest = pointer[len('definitions/'):]
    parts = rest.split('/')
    if not parts:
        return None
    fd = flat_defs.get(file_abs, {})
    head = parts[0]
    if head not in fd:
        return None
    obj = fd[head]
    for seg in parts[1:]:
        seg = seg.replace('~1', '/').replace('~0', '~')
        if isinstance(obj, dict) and seg in obj:
            obj = obj[seg]
        elif isinstance(obj, list) and seg.isdigit() and int(seg) < len(obj):
            obj = obj[int(seg)]
        else:
            return None
    return obj

def resolve_ref(file_abs, ref):
    if ref.startswith('#'):
        pointer = ref[1:].lstrip('/')
        if pointer.startswith('definitions/'):
            return (file_abs, pointer)
        return None
    if '.json' in ref:
        fpath, _, frag = ref.partition('#')
        fpath = canon(os.path.join(os.path.dirname(file_abs), fpath))
        if fpath not in files:
            return None
        pointer = frag.lstrip('#/')
        if pointer.startswith('definitions/'):
            return (fpath, pointer)
    return None

def comp_name_for(file_abs, pointer):
    rest = pointer[len('definitions'):].strip('/')
    if '/' in rest:
        return f"{path_key(file_abs)}_{rest.replace('/', '_')}"
    return map_name(file_abs, rest)

def add_component(file_abs, pointer):
    key = (file_abs, pointer)
    if key in FIX_KEYS:
        val = FIX_KEYS[key]
        if isinstance(val, tuple) and val[0] == 'POINTER':
            return add_component(canon(os.path.join(ROOT, val[1])), val[2])
        # inline schema fix: create a unique component holding the literal schema
        rest = pointer[len('definitions'):].strip('/')
        cname = f"{path_key(file_abs)}_{rest.replace('/', '_')}"
        cname = cname.replace('.', '_')
        if cname not in components:
            components[cname] = val
            pending.append((file_abs, cname, val, cname))
        return cname
    schema = get_pointer_schema(file_abs, pointer)
    if schema is None:
        print(f"!! POINTER MISS: {os.path.relpath(file_abs, ROOT)} #{pointer}")
        return None
    cname = comp_name_for(file_abs, pointer)
    if cname in components:
        return cname
    components[cname] = schema
    pending.append((file_abs, cname))
    return cname

pending = []

# ---------- Seed from main spec ----------
main = yaml.safe_load(open(MAIN))
MAIN_ABS = canon(MAIN)

def rewrite_ref_ctx(obj, file_abs):
    if isinstance(obj, dict):
        if '$ref' in obj:
            ref = obj['$ref']
            if isinstance(ref, str):
                res = resolve_ref(file_abs, ref)
                if res:
                    tfile, tpointer = res
                    tcname = add_component(tfile, tpointer)
                    if tcname:
                        obj['$ref'] = f"#/components/schemas/{tcname}"
            # always continue into sibling keys (oneOf/anyOf/allOf next to $ref)
        for k, v in list(obj.items()):
            if k == '$ref':
                continue
            rewrite_ref_ctx(v, file_abs)
    elif isinstance(obj, list):
        for v in obj:
            rewrite_ref_ctx(v, file_abs)

rewrite_ref_ctx(main, MAIN_ABS)

# Worklist
processed = 0
while pending:
    item = pending.pop()
    file_abs, cname = item[0], item[1]
    if len(item) == 3:
        continue  # inline fix already placed
    schema = components[cname]
    # note: pending items were queued by add_component; the (file_abs, cname) is the source context
    rewrite_ref_ctx(schema, file_abs)
    processed += 1

print("Components included:", len(components))
print("Processed entries:", processed)

# ---------- Normalize + Deduplicate operationIds ----------
import re as _re
for path, item in (main.get('paths') or {}).items():
    for method, op in item.items():
        if not isinstance(op, dict) or 'operationId' not in op:
            continue
        oid = op['operationId']
        nid = _re.sub(r'\s+', '_', oid.strip())
        nid = _re.sub(r'[^A-Za-z0-9_.]', '_', nid)
        if nid != oid:
            print(f"NORMALIZE {method.upper()} {path}: {oid!r} -> {nid!r}")
            op['operationId'] = nid

def _to_snake(name):
    s = _re.sub(r'(?<=[a-z0-9])(?=[A-Z])', '_', name)
    s = _re.sub(r'[^A-Za-z0-9]+', '_', s).strip('_').lower()
    return s

seen_snakes = set()
for path, item in (main.get('paths') or {}).items():
    for method, op in item.items():
        if not isinstance(op, dict) or 'operationId' not in op:
            continue
        oid = op['operationId']
        snake = _to_snake(oid)
        if snake in seen_snakes:
            suffix = path.strip('/').replace('/', '_').replace('-', '_').replace('{', '').replace('}', '')
            new_oid = f"{oid}_{suffix}"
            i = 2
            while _to_snake(new_oid) in seen_snakes:
                new_oid = f"{oid}_{suffix}_{i}"
                i += 1
            print(f"DEDUPE {method.upper()} {path}: {oid} -> {new_oid}")
            op['operationId'] = new_oid
        seen_snakes.add(_to_snake(op['operationId']))

# Merge into main spec
main.setdefault('components', {})
main['components'].setdefault('schemas', {})
for name, schema in components.items():
    if name not in main['components']['schemas']:
        main['components']['schemas'][name] = schema

def strip_noise(obj):
    if isinstance(obj, dict):
        for k in ('id', '$schema', '$comment'):
            if k in obj:
                del obj[k]
        for v in obj.values():
            strip_noise(v)
    elif isinstance(obj, list):
        for v in obj:
            strip_noise(v)

strip_noise(main)

with open(os.path.join(ROOT, 'openapi.fm.bundled.yaml'), 'w') as f:
    yaml.safe_dump(main, f, default_flow_style=False, sort_keys=False, width=10000)
print("Bundled YAML written")
