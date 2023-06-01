





ROTR[n_, m_Integer] := BitShiftRight[n, m] + BitShiftLeft[Mod[n, 2^m], (32 - m)]

s0[a_] := BitXor[ROTR[a,7], ROTR[a,18], BitShiftRight[a,3] ]

s1[a_] := BitXor[ROTR[a,17], ROTR[a,19], BitShiftRight[a,10] ]

E0[a_] := BitXor[ROTR[a,2], ROTR[a,13], ROTR[a,22] ]

E1[a_] := BitXor[ROTR[a,6], ROTR[a,11], ROTR[a,25] ]

Ch[a_,b_,c_] := BitXor[BitAnd[a,b], BitAnd[BitNot[a],c] ] 

Maj[a_,b_,c_] := BitXor[BitAnd[a,b], BitAnd[a,c], BitAnd[c,b] ]


PrintBase2[x_Integer] := Print[ToString[BaseForm[x,2] ] ]



a = 8
b = 2

Print[s1[a] ]



