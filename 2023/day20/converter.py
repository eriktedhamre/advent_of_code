input_text = """
#mq -> zt
#jr -> gt, hv
#qh -> gt, kt
#hr -> hg, zt
#px -> nl
#fx -> mv
#tz -> jl, zt
#mv -> xd, ss
#cb -> sj, zt
#sn -> kx
#xp -> vl, zt
#nl -> hz
#dp -> bj, xd
#zq -> xd, fx
#hv -> gt
#zm -> ms, vr
&ct -> bb
&xd -> kp, bg, ss, sn, mf, qb, fx
&kp -> bb
&gt -> pm, xh, gp, nn, bv, ct
#ss -> bg
&zt -> jl, xc, jf, fh
&ms -> br, nl, px, vg, vr, ks, fr
#xj -> ms, vt
#ts -> ms
#lt -> gt, xh
#gp -> bx
#br -> px
#sj -> mq, zt
broadcaster -> fr, jf, mf, bv
#jl -> cb
#mf -> xd, qb
#vl -> zt, tz
&ks -> bb
&bb -> rx
#bv -> gt, nn
#bs -> xj, ms
#vt -> ms, ts
#nn -> nz
#nz -> pm, gt
#xh -> qh
#xl -> xd, sn
#fr -> ms, zm
#pd -> hr, zt
#pm -> lt
#vg -> bs
#bj -> xd
#fh -> xp
#qb -> zq
#kx -> dp, xd
#bx -> jr, gt
#vr -> br
#hg -> fh, zt
#kt -> gp, gt
#hz -> ms, vg
#jf -> zt, pd
#bg -> xl
&xc -> bb
"""

# Parse the input and keep prefixes for both sides
edges = []
node_type = {}
for line in input_text.strip().splitlines():
    if "->" not in line:
        continue
    src, dests = line.split("->", 1)
    src = src.strip()
    node_type[src[1:]] = src[0]
    for d in dests.split(","):
        dst = d.strip()
        edges.append((src, dst))
    


# Write DOT file
with open("graph.dot", "w") as f:
    f.write('digraph G {\n')
    
    
    for src, dst in edges:
        dst_str = dst
        if dst == "broadcaster" or dst == "rx":
            print(dst)
        elif node_type[dst] == "#":
            dst_str = "#"+dst
        elif node_type[dst] == "&":
            dst_str = "&"+dst
        f.write(f'    "{src}" -> "{dst_str}";\n')
    f.write("\n")
    f.write("}\n")

print("✅ Wrote graph.dot")