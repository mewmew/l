// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io/ioutil"
	"unicode/utf8"

	"github.com/mewmew/l/asm/internal/token"
)

const (
	NoState    = -1
	NumStates  = 2693
	NumSymbols = 3477
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (l *Lexer) Scan() (tok *token.Token) {
	tok = new(token.Token)
	if l.pos >= len(l.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = l.pos, l.line, l.column
		return
	}
	start, startLine, startColumn, end := l.pos, l.line, l.column, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
		if l.pos >= len(l.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(l.src[l.pos:])
			l.pos += size
		}

		nextState := -1
		if rune1 != -1 {
			nextState = TransTab[state](rune1)
		}
		state = nextState

		if state != -1 {

			switch rune1 {
			case '\n':
				l.line++
				l.column = 1
			case '\r':
				l.column = 1
			case '\t':
				l.column += 4
			default:
				l.column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				end = l.pos
			case ActTab[state].Ignore != "":
				start, startLine, startColumn = l.pos, l.line, l.column
				state = 0
				if start >= len(l.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = l.pos
			}
		}
	}
	if end > start {
		l.pos = end
		tok.Lit = l.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = start, startLine, startColumn

	return
}

func (l *Lexer) Reset() {
	l.pos = 0
}

/*
Lexer symbols:
0: ':'
1: ':'
2: '#'
3: '$'
4: '!'
5: '!'
6: 'D'
7: 'W'
8: '_'
9: 'T'
10: 'A'
11: 'G'
12: '_'
13: '_'
14: 'D'
15: 'W'
16: '_'
17: 'A'
18: 'T'
19: 'E'
20: '_'
21: '_'
22: 'D'
23: 'I'
24: 'F'
25: 'l'
26: 'a'
27: 'g'
28: '_'
29: 'D'
30: 'W'
31: '_'
32: 'L'
33: 'A'
34: 'N'
35: 'G'
36: '_'
37: '_'
38: 'D'
39: 'W'
40: '_'
41: 'C'
42: 'C'
43: '_'
44: '_'
45: 'C'
46: 'S'
47: 'K'
48: '_'
49: '_'
50: 'D'
51: 'W'
52: '_'
53: 'V'
54: 'I'
55: 'R'
56: 'T'
57: 'U'
58: 'A'
59: 'L'
60: 'I'
61: 'T'
62: 'Y'
63: '_'
64: '_'
65: 'D'
66: 'W'
67: '_'
68: 'M'
69: 'A'
70: 'C'
71: 'I'
72: 'N'
73: 'F'
74: 'O'
75: '_'
76: '_'
77: 'D'
78: 'W'
79: '_'
80: 'O'
81: 'P'
82: '_'
83: '_'
84: 'i'
85: 's'
86: 'o'
87: 'u'
88: 'r'
89: 'c'
90: 'e'
91: '_'
92: 'f'
93: 'i'
94: 'l'
95: 'e'
96: 'n'
97: 'a'
98: 'm'
99: 'e'
100: '='
101: 't'
102: 'a'
103: 'r'
104: 'g'
105: 'e'
106: 't'
107: 'd'
108: 'a'
109: 't'
110: 'a'
111: 'l'
112: 'a'
113: 'y'
114: 'o'
115: 'u'
116: 't'
117: 't'
118: 'r'
119: 'i'
120: 'p'
121: 'l'
122: 'e'
123: 'm'
124: 'o'
125: 'd'
126: 'u'
127: 'l'
128: 'e'
129: 'a'
130: 's'
131: 'm'
132: 't'
133: 'y'
134: 'p'
135: 'e'
136: 'c'
137: 'o'
138: 'm'
139: 'd'
140: 'a'
141: 't'
142: 'a'
143: 'n'
144: 'y'
145: 'e'
146: 'x'
147: 'a'
148: 'c'
149: 't'
150: 'm'
151: 'a'
152: 't'
153: 'c'
154: 'h'
155: 'l'
156: 'a'
157: 'r'
158: 'g'
159: 'e'
160: 's'
161: 't'
162: 'n'
163: 'o'
164: 'd'
165: 'u'
166: 'p'
167: 'l'
168: 'i'
169: 'c'
170: 'a'
171: 't'
172: 'e'
173: 's'
174: 's'
175: 'a'
176: 'm'
177: 'e'
178: 's'
179: 'i'
180: 'z'
181: 'e'
182: 'e'
183: 'x'
184: 't'
185: 'e'
186: 'r'
187: 'n'
188: 'a'
189: 'l'
190: 'l'
191: 'y'
192: '_'
193: 'i'
194: 'n'
195: 'i'
196: 't'
197: 'i'
198: 'a'
199: 'l'
200: 'i'
201: 'z'
202: 'e'
203: 'd'
204: 'c'
205: 'o'
206: 'n'
207: 's'
208: 't'
209: 'a'
210: 'n'
211: 't'
212: 'g'
213: 'l'
214: 'o'
215: 'b'
216: 'a'
217: 'l'
218: ','
219: 'a'
220: 'l'
221: 'i'
222: 'a'
223: 's'
224: 'i'
225: 'f'
226: 'u'
227: 'n'
228: 'c'
229: 'd'
230: 'e'
231: 'c'
232: 'l'
233: 'a'
234: 'r'
235: 'e'
236: 'd'
237: 'e'
238: 'f'
239: 'i'
240: 'n'
241: 'e'
242: '('
243: ')'
244: 'g'
245: 'c'
246: 'p'
247: 'r'
248: 'e'
249: 'f'
250: 'i'
251: 'x'
252: 'p'
253: 'r'
254: 'o'
255: 'l'
256: 'o'
257: 'g'
258: 'u'
259: 'e'
260: 'p'
261: 'e'
262: 'r'
263: 's'
264: 'o'
265: 'n'
266: 'a'
267: 'l'
268: 'i'
269: 't'
270: 'y'
271: '{'
272: '}'
273: 'a'
274: 't'
275: 't'
276: 'r'
277: 'i'
278: 'b'
279: 'u'
280: 't'
281: 'e'
282: 's'
283: '!'
284: 'd'
285: 'i'
286: 's'
287: 't'
288: 'i'
289: 'n'
290: 'c'
291: 't'
292: 'u'
293: 's'
294: 'e'
295: 'l'
296: 'i'
297: 's'
298: 't'
299: 'o'
300: 'r'
301: 'd'
302: 'e'
303: 'r'
304: 'u'
305: 's'
306: 'e'
307: 'l'
308: 'i'
309: 's'
310: 't'
311: 'o'
312: 'r'
313: 'd'
314: 'e'
315: 'r'
316: '_'
317: 'b'
318: 'b'
319: 'v'
320: 'o'
321: 'i'
322: 'd'
323: 'h'
324: 'a'
325: 'l'
326: 'f'
327: 'f'
328: 'l'
329: 'o'
330: 'a'
331: 't'
332: 'd'
333: 'o'
334: 'u'
335: 'b'
336: 'l'
337: 'e'
338: 'x'
339: '8'
340: '6'
341: '_'
342: 'f'
343: 'p'
344: '8'
345: '0'
346: 'f'
347: 'p'
348: '1'
349: '2'
350: '8'
351: 'p'
352: 'p'
353: 'c'
354: '_'
355: 'f'
356: 'p'
357: '1'
358: '2'
359: '8'
360: 'x'
361: '8'
362: '6'
363: '_'
364: 'm'
365: 'm'
366: 'x'
367: '*'
368: 'a'
369: 'd'
370: 'd'
371: 'r'
372: 's'
373: 'p'
374: 'a'
375: 'c'
376: 'e'
377: '<'
378: 'x'
379: '>'
380: 'l'
381: 'a'
382: 'b'
383: 'e'
384: 'l'
385: 't'
386: 'o'
387: 'k'
388: 'e'
389: 'n'
390: 'm'
391: 'e'
392: 't'
393: 'a'
394: 'd'
395: 'a'
396: 't'
397: 'a'
398: '['
399: ']'
400: 'o'
401: 'p'
402: 'a'
403: 'q'
404: 'u'
405: 'e'
406: 's'
407: 'i'
408: 'd'
409: 'e'
410: 'e'
411: 'f'
412: 'f'
413: 'e'
414: 'c'
415: 't'
416: 'a'
417: 'l'
418: 'i'
419: 'g'
420: 'n'
421: 's'
422: 't'
423: 'a'
424: 'c'
425: 'k'
426: 'i'
427: 'n'
428: 't'
429: 'e'
430: 'l'
431: 'd'
432: 'i'
433: 'a'
434: 'l'
435: 'e'
436: 'c'
437: 't'
438: 't'
439: 'r'
440: 'u'
441: 'e'
442: 'f'
443: 'a'
444: 'l'
445: 's'
446: 'e'
447: 'n'
448: 'u'
449: 'l'
450: 'l'
451: 'n'
452: 'o'
453: 'n'
454: 'e'
455: 'c'
456: 'z'
457: 'e'
458: 'r'
459: 'o'
460: 'i'
461: 'n'
462: 'i'
463: 't'
464: 'i'
465: 'a'
466: 'l'
467: 'i'
468: 'z'
469: 'e'
470: 'r'
471: 'u'
472: 'n'
473: 'd'
474: 'e'
475: 'f'
476: 'b'
477: 'l'
478: 'o'
479: 'c'
480: 'k'
481: 'a'
482: 'd'
483: 'd'
484: 'r'
485: 'e'
486: 's'
487: 's'
488: 'a'
489: 'd'
490: 'd'
491: 'f'
492: 'a'
493: 'd'
494: 'd'
495: 's'
496: 'u'
497: 'b'
498: 'f'
499: 's'
500: 'u'
501: 'b'
502: 'm'
503: 'u'
504: 'l'
505: 'f'
506: 'm'
507: 'u'
508: 'l'
509: 'u'
510: 'd'
511: 'i'
512: 'v'
513: 's'
514: 'd'
515: 'i'
516: 'v'
517: 'f'
518: 'd'
519: 'i'
520: 'v'
521: 'u'
522: 'r'
523: 'e'
524: 'm'
525: 's'
526: 'r'
527: 'e'
528: 'm'
529: 'f'
530: 'r'
531: 'e'
532: 'm'
533: 's'
534: 'h'
535: 'l'
536: 'l'
537: 's'
538: 'h'
539: 'r'
540: 'a'
541: 's'
542: 'h'
543: 'r'
544: 'a'
545: 'n'
546: 'd'
547: 'o'
548: 'r'
549: 'x'
550: 'o'
551: 'r'
552: 'e'
553: 'x'
554: 't'
555: 'r'
556: 'a'
557: 'c'
558: 't'
559: 'e'
560: 'l'
561: 'e'
562: 'm'
563: 'e'
564: 'n'
565: 't'
566: 'i'
567: 'n'
568: 's'
569: 'e'
570: 'r'
571: 't'
572: 'e'
573: 'l'
574: 'e'
575: 'm'
576: 'e'
577: 'n'
578: 't'
579: 's'
580: 'h'
581: 'u'
582: 'f'
583: 'f'
584: 'l'
585: 'e'
586: 'v'
587: 'e'
588: 'c'
589: 't'
590: 'o'
591: 'r'
592: 'e'
593: 'x'
594: 't'
595: 'r'
596: 'a'
597: 'c'
598: 't'
599: 'v'
600: 'a'
601: 'l'
602: 'u'
603: 'e'
604: 'i'
605: 'n'
606: 's'
607: 'e'
608: 'r'
609: 't'
610: 'v'
611: 'a'
612: 'l'
613: 'u'
614: 'e'
615: 'g'
616: 'e'
617: 't'
618: 'e'
619: 'l'
620: 'e'
621: 'm'
622: 'e'
623: 'n'
624: 't'
625: 'p'
626: 't'
627: 'r'
628: 'i'
629: 'n'
630: 'r'
631: 'a'
632: 'n'
633: 'g'
634: 'e'
635: 't'
636: 'r'
637: 'u'
638: 'n'
639: 'c'
640: 't'
641: 'o'
642: 'z'
643: 'e'
644: 'x'
645: 't'
646: 's'
647: 'e'
648: 'x'
649: 't'
650: 'f'
651: 'p'
652: 't'
653: 'r'
654: 'u'
655: 'n'
656: 'c'
657: 'f'
658: 'p'
659: 'e'
660: 'x'
661: 't'
662: 'f'
663: 'p'
664: 't'
665: 'o'
666: 'u'
667: 'i'
668: 'f'
669: 'p'
670: 't'
671: 'o'
672: 's'
673: 'i'
674: 'u'
675: 'i'
676: 't'
677: 'o'
678: 'f'
679: 'p'
680: 's'
681: 'i'
682: 't'
683: 'o'
684: 'f'
685: 'p'
686: 'p'
687: 't'
688: 'r'
689: 't'
690: 'o'
691: 'i'
692: 'n'
693: 't'
694: 'i'
695: 'n'
696: 't'
697: 't'
698: 'o'
699: 'p'
700: 't'
701: 'r'
702: 'b'
703: 'i'
704: 't'
705: 'c'
706: 'a'
707: 's'
708: 't'
709: 'a'
710: 'd'
711: 'd'
712: 'r'
713: 's'
714: 'p'
715: 'a'
716: 'c'
717: 'e'
718: 'c'
719: 'a'
720: 's'
721: 't'
722: 'i'
723: 'c'
724: 'm'
725: 'p'
726: 'f'
727: 'c'
728: 'm'
729: 'p'
730: 's'
731: 'e'
732: 'l'
733: 'e'
734: 'c'
735: 't'
736: 'a'
737: 'l'
738: 'l'
739: 'o'
740: 'c'
741: 'a'
742: 'i'
743: 'n'
744: 'a'
745: 'l'
746: 'l'
747: 'o'
748: 'c'
749: 'a'
750: 's'
751: 'w'
752: 'i'
753: 'f'
754: 't'
755: 'e'
756: 'r'
757: 'r'
758: 'o'
759: 'r'
760: 'l'
761: 'o'
762: 'a'
763: 'd'
764: 'a'
765: 't'
766: 'o'
767: 'm'
768: 'i'
769: 'c'
770: 's'
771: 't'
772: 'o'
773: 'r'
774: 'e'
775: 'f'
776: 'e'
777: 'n'
778: 'c'
779: 'e'
780: 'c'
781: 'm'
782: 'p'
783: 'x'
784: 'c'
785: 'h'
786: 'g'
787: 'w'
788: 'e'
789: 'a'
790: 'k'
791: 'a'
792: 't'
793: 'o'
794: 'm'
795: 'i'
796: 'c'
797: 'r'
798: 'm'
799: 'w'
800: 'm'
801: 'a'
802: 'x'
803: 'm'
804: 'i'
805: 'n'
806: 'n'
807: 'a'
808: 'n'
809: 'd'
810: 'u'
811: 'm'
812: 'a'
813: 'x'
814: 'u'
815: 'm'
816: 'i'
817: 'n'
818: 'x'
819: 'c'
820: 'h'
821: 'g'
822: 'p'
823: 'h'
824: 'i'
825: 'c'
826: 'a'
827: 'l'
828: 'l'
829: 'm'
830: 'u'
831: 's'
832: 't'
833: 't'
834: 'a'
835: 'i'
836: 'l'
837: 'n'
838: 'o'
839: 't'
840: 'a'
841: 'i'
842: 'l'
843: 't'
844: 'a'
845: 'i'
846: 'l'
847: 'v'
848: 'a'
849: '_'
850: 'a'
851: 'r'
852: 'g'
853: 'l'
854: 'a'
855: 'n'
856: 'd'
857: 'i'
858: 'n'
859: 'g'
860: 'p'
861: 'a'
862: 'd'
863: 'c'
864: 'l'
865: 'e'
866: 'a'
867: 'n'
868: 'u'
869: 'p'
870: 'c'
871: 'a'
872: 't'
873: 'c'
874: 'h'
875: 'f'
876: 'i'
877: 'l'
878: 't'
879: 'e'
880: 'r'
881: 'c'
882: 'a'
883: 't'
884: 'c'
885: 'h'
886: 'p'
887: 'a'
888: 'd'
889: 'w'
890: 'i'
891: 't'
892: 'h'
893: 'i'
894: 'n'
895: 'c'
896: 'l'
897: 'e'
898: 'a'
899: 'n'
900: 'u'
901: 'p'
902: 'p'
903: 'a'
904: 'd'
905: 'r'
906: 'e'
907: 't'
908: 'b'
909: 'r'
910: 's'
911: 'w'
912: 'i'
913: 't'
914: 'c'
915: 'h'
916: 'i'
917: 'n'
918: 'd'
919: 'i'
920: 'r'
921: 'e'
922: 'c'
923: 't'
924: 'b'
925: 'r'
926: 'i'
927: 'n'
928: 'v'
929: 'o'
930: 'k'
931: 'e'
932: 'u'
933: 'n'
934: 'w'
935: 'i'
936: 'n'
937: 'd'
938: 'r'
939: 'e'
940: 's'
941: 'u'
942: 'm'
943: 'e'
944: 'c'
945: 'a'
946: 't'
947: 'c'
948: 'h'
949: 's'
950: 'w'
951: 'i'
952: 't'
953: 'c'
954: 'h'
955: 'c'
956: 'a'
957: 't'
958: 'c'
959: 'h'
960: 'r'
961: 'e'
962: 't'
963: 'f'
964: 'r'
965: 'o'
966: 'm'
967: 'c'
968: 'l'
969: 'e'
970: 'a'
971: 'n'
972: 'u'
973: 'p'
974: 'r'
975: 'e'
976: 't'
977: 'u'
978: 'n'
979: 'r'
980: 'e'
981: 'a'
982: 'c'
983: 'h'
984: 'a'
985: 'b'
986: 'l'
987: 'e'
988: 'c'
989: 'a'
990: 'l'
991: 'l'
992: 'e'
993: 'r'
994: '!'
995: 'D'
996: 'I'
997: 'C'
998: 'o'
999: 'm'
1000: 'p'
1001: 'i'
1002: 'l'
1003: 'e'
1004: 'U'
1005: 'n'
1006: 'i'
1007: 't'
1008: 'l'
1009: 'a'
1010: 'n'
1011: 'g'
1012: 'u'
1013: 'a'
1014: 'g'
1015: 'e'
1016: ':'
1017: 'p'
1018: 'r'
1019: 'o'
1020: 'd'
1021: 'u'
1022: 'c'
1023: 'e'
1024: 'r'
1025: ':'
1026: 'f'
1027: 'l'
1028: 'a'
1029: 'g'
1030: 's'
1031: ':'
1032: 'r'
1033: 'u'
1034: 'n'
1035: 't'
1036: 'i'
1037: 'm'
1038: 'e'
1039: 'V'
1040: 'e'
1041: 'r'
1042: 's'
1043: 'i'
1044: 'o'
1045: 'n'
1046: ':'
1047: 's'
1048: 'p'
1049: 'l'
1050: 'i'
1051: 't'
1052: 'D'
1053: 'e'
1054: 'b'
1055: 'u'
1056: 'g'
1057: 'F'
1058: 'i'
1059: 'l'
1060: 'e'
1061: 'n'
1062: 'a'
1063: 'm'
1064: 'e'
1065: ':'
1066: 'e'
1067: 'm'
1068: 'i'
1069: 's'
1070: 's'
1071: 'i'
1072: 'o'
1073: 'n'
1074: 'K'
1075: 'i'
1076: 'n'
1077: 'd'
1078: ':'
1079: 'e'
1080: 'n'
1081: 'u'
1082: 'm'
1083: 's'
1084: ':'
1085: 'r'
1086: 'e'
1087: 't'
1088: 'a'
1089: 'i'
1090: 'n'
1091: 'e'
1092: 'd'
1093: 'T'
1094: 'y'
1095: 'p'
1096: 'e'
1097: 's'
1098: ':'
1099: 'g'
1100: 'l'
1101: 'o'
1102: 'b'
1103: 'a'
1104: 'l'
1105: 's'
1106: ':'
1107: 'i'
1108: 'm'
1109: 'p'
1110: 'o'
1111: 'r'
1112: 't'
1113: 's'
1114: ':'
1115: 'm'
1116: 'a'
1117: 'c'
1118: 'r'
1119: 'o'
1120: 's'
1121: ':'
1122: 'd'
1123: 'w'
1124: 'o'
1125: 'I'
1126: 'd'
1127: ':'
1128: 's'
1129: 'p'
1130: 'l'
1131: 'i'
1132: 't'
1133: 'D'
1134: 'e'
1135: 'b'
1136: 'u'
1137: 'g'
1138: 'I'
1139: 'n'
1140: 'l'
1141: 'i'
1142: 'n'
1143: 'i'
1144: 'n'
1145: 'g'
1146: ':'
1147: 'd'
1148: 'e'
1149: 'b'
1150: 'u'
1151: 'g'
1152: 'I'
1153: 'n'
1154: 'f'
1155: 'o'
1156: 'F'
1157: 'o'
1158: 'r'
1159: 'P'
1160: 'r'
1161: 'o'
1162: 'f'
1163: 'i'
1164: 'l'
1165: 'i'
1166: 'n'
1167: 'g'
1168: ':'
1169: 'g'
1170: 'n'
1171: 'u'
1172: 'P'
1173: 'u'
1174: 'b'
1175: 'n'
1176: 'a'
1177: 'm'
1178: 'e'
1179: 's'
1180: ':'
1181: '!'
1182: 'D'
1183: 'I'
1184: 'F'
1185: 'i'
1186: 'l'
1187: 'e'
1188: 'f'
1189: 'i'
1190: 'l'
1191: 'e'
1192: 'n'
1193: 'a'
1194: 'm'
1195: 'e'
1196: ':'
1197: 'd'
1198: 'i'
1199: 'r'
1200: 'e'
1201: 'c'
1202: 't'
1203: 'o'
1204: 'r'
1205: 'y'
1206: ':'
1207: 'c'
1208: 'h'
1209: 'e'
1210: 'c'
1211: 'k'
1212: 's'
1213: 'u'
1214: 'm'
1215: 'k'
1216: 'i'
1217: 'n'
1218: 'd'
1219: ':'
1220: 'c'
1221: 'h'
1222: 'e'
1223: 'c'
1224: 'k'
1225: 's'
1226: 'u'
1227: 'm'
1228: ':'
1229: '!'
1230: 'D'
1231: 'I'
1232: 'B'
1233: 'a'
1234: 's'
1235: 'i'
1236: 'c'
1237: 'T'
1238: 'y'
1239: 'p'
1240: 'e'
1241: 'e'
1242: 'n'
1243: 'c'
1244: 'o'
1245: 'd'
1246: 'i'
1247: 'n'
1248: 'g'
1249: ':'
1250: '!'
1251: 'D'
1252: 'I'
1253: 'S'
1254: 'u'
1255: 'b'
1256: 'r'
1257: 'o'
1258: 'u'
1259: 't'
1260: 'i'
1261: 'n'
1262: 'e'
1263: 'T'
1264: 'y'
1265: 'p'
1266: 'e'
1267: 'c'
1268: 'c'
1269: ':'
1270: 't'
1271: 'y'
1272: 'p'
1273: 'e'
1274: 's'
1275: ':'
1276: '!'
1277: 'D'
1278: 'I'
1279: 'D'
1280: 'e'
1281: 'r'
1282: 'i'
1283: 'v'
1284: 'e'
1285: 'd'
1286: 'T'
1287: 'y'
1288: 'p'
1289: 'e'
1290: 'e'
1291: 'x'
1292: 't'
1293: 'r'
1294: 'a'
1295: 'D'
1296: 'a'
1297: 't'
1298: 'a'
1299: ':'
1300: 'd'
1301: 'w'
1302: 'a'
1303: 'r'
1304: 'f'
1305: 'A'
1306: 'd'
1307: 'd'
1308: 'r'
1309: 'e'
1310: 's'
1311: 's'
1312: 'S'
1313: 'p'
1314: 'a'
1315: 'c'
1316: 'e'
1317: ':'
1318: '!'
1319: 'D'
1320: 'I'
1321: 'C'
1322: 'o'
1323: 'm'
1324: 'p'
1325: 'o'
1326: 's'
1327: 'i'
1328: 't'
1329: 'e'
1330: 'T'
1331: 'y'
1332: 'p'
1333: 'e'
1334: 'e'
1335: 'l'
1336: 'e'
1337: 'm'
1338: 'e'
1339: 'n'
1340: 't'
1341: 's'
1342: ':'
1343: 'r'
1344: 'u'
1345: 'n'
1346: 't'
1347: 'i'
1348: 'm'
1349: 'e'
1350: 'L'
1351: 'a'
1352: 'n'
1353: 'g'
1354: ':'
1355: 'v'
1356: 't'
1357: 'a'
1358: 'b'
1359: 'l'
1360: 'e'
1361: 'H'
1362: 'o'
1363: 'l'
1364: 'd'
1365: 'e'
1366: 'r'
1367: ':'
1368: 'i'
1369: 'd'
1370: 'e'
1371: 'n'
1372: 't'
1373: 'i'
1374: 'f'
1375: 'i'
1376: 'e'
1377: 'r'
1378: ':'
1379: 'd'
1380: 'i'
1381: 's'
1382: 'c'
1383: 'r'
1384: 'i'
1385: 'm'
1386: 'i'
1387: 'n'
1388: 'a'
1389: 't'
1390: 'o'
1391: 'r'
1392: ':'
1393: '!'
1394: 'D'
1395: 'I'
1396: 'S'
1397: 'u'
1398: 'b'
1399: 'r'
1400: 'a'
1401: 'n'
1402: 'g'
1403: 'e'
1404: 'c'
1405: 'o'
1406: 'u'
1407: 'n'
1408: 't'
1409: ':'
1410: 'l'
1411: 'o'
1412: 'w'
1413: 'e'
1414: 'r'
1415: 'B'
1416: 'o'
1417: 'u'
1418: 'n'
1419: 'd'
1420: ':'
1421: '!'
1422: 'D'
1423: 'I'
1424: 'E'
1425: 'n'
1426: 'u'
1427: 'm'
1428: 'e'
1429: 'r'
1430: 'a'
1431: 't'
1432: 'o'
1433: 'r'
1434: 'v'
1435: 'a'
1436: 'l'
1437: 'u'
1438: 'e'
1439: ':'
1440: 'i'
1441: 's'
1442: 'U'
1443: 'n'
1444: 's'
1445: 'i'
1446: 'g'
1447: 'n'
1448: 'e'
1449: 'd'
1450: ':'
1451: '!'
1452: 'D'
1453: 'I'
1454: 'T'
1455: 'e'
1456: 'm'
1457: 'p'
1458: 'l'
1459: 'a'
1460: 't'
1461: 'e'
1462: 'T'
1463: 'y'
1464: 'p'
1465: 'e'
1466: 'P'
1467: 'a'
1468: 'r'
1469: 'a'
1470: 'm'
1471: 'e'
1472: 't'
1473: 'e'
1474: 'r'
1475: '!'
1476: 'D'
1477: 'I'
1478: 'T'
1479: 'e'
1480: 'm'
1481: 'p'
1482: 'l'
1483: 'a'
1484: 't'
1485: 'e'
1486: 'V'
1487: 'a'
1488: 'l'
1489: 'u'
1490: 'e'
1491: 'P'
1492: 'a'
1493: 'r'
1494: 'a'
1495: 'm'
1496: 'e'
1497: 't'
1498: 'e'
1499: 'r'
1500: '!'
1501: 'D'
1502: 'I'
1503: 'M'
1504: 'o'
1505: 'd'
1506: 'u'
1507: 'l'
1508: 'e'
1509: 'c'
1510: 'o'
1511: 'n'
1512: 'f'
1513: 'i'
1514: 'g'
1515: 'M'
1516: 'a'
1517: 'c'
1518: 'r'
1519: 'o'
1520: 's'
1521: ':'
1522: 'i'
1523: 'n'
1524: 'c'
1525: 'l'
1526: 'u'
1527: 'd'
1528: 'e'
1529: 'P'
1530: 'a'
1531: 't'
1532: 'h'
1533: ':'
1534: 'i'
1535: 's'
1536: 'y'
1537: 's'
1538: 'r'
1539: 'o'
1540: 'o'
1541: 't'
1542: ':'
1543: '!'
1544: 'D'
1545: 'I'
1546: 'N'
1547: 'a'
1548: 'm'
1549: 'e'
1550: 's'
1551: 'p'
1552: 'a'
1553: 'c'
1554: 'e'
1555: 'e'
1556: 'x'
1557: 'p'
1558: 'o'
1559: 'r'
1560: 't'
1561: 'S'
1562: 'y'
1563: 'm'
1564: 'b'
1565: 'o'
1566: 'l'
1567: 's'
1568: ':'
1569: '!'
1570: 'D'
1571: 'I'
1572: 'G'
1573: 'l'
1574: 'o'
1575: 'b'
1576: 'a'
1577: 'l'
1578: 'V'
1579: 'a'
1580: 'r'
1581: 'i'
1582: 'a'
1583: 'b'
1584: 'l'
1585: 'e'
1586: '!'
1587: 'D'
1588: 'I'
1589: 'S'
1590: 'u'
1591: 'b'
1592: 'p'
1593: 'r'
1594: 'o'
1595: 'g'
1596: 'r'
1597: 'a'
1598: 'm'
1599: 's'
1600: 'c'
1601: 'o'
1602: 'p'
1603: 'e'
1604: 'L'
1605: 'i'
1606: 'n'
1607: 'e'
1608: ':'
1609: 'c'
1610: 'o'
1611: 'n'
1612: 't'
1613: 'a'
1614: 'i'
1615: 'n'
1616: 'i'
1617: 'n'
1618: 'g'
1619: 'T'
1620: 'y'
1621: 'p'
1622: 'e'
1623: ':'
1624: 'v'
1625: 'i'
1626: 'r'
1627: 't'
1628: 'u'
1629: 'a'
1630: 'l'
1631: 'i'
1632: 't'
1633: 'y'
1634: ':'
1635: 'v'
1636: 'i'
1637: 'r'
1638: 't'
1639: 'u'
1640: 'a'
1641: 'l'
1642: 'I'
1643: 'n'
1644: 'd'
1645: 'e'
1646: 'x'
1647: ':'
1648: 't'
1649: 'h'
1650: 'i'
1651: 's'
1652: 'A'
1653: 'd'
1654: 'j'
1655: 'u'
1656: 's'
1657: 't'
1658: 'm'
1659: 'e'
1660: 'n'
1661: 't'
1662: ':'
1663: 'u'
1664: 'n'
1665: 'i'
1666: 't'
1667: ':'
1668: 'v'
1669: 'a'
1670: 'r'
1671: 'i'
1672: 'a'
1673: 'b'
1674: 'l'
1675: 'e'
1676: 's'
1677: ':'
1678: 't'
1679: 'h'
1680: 'r'
1681: 'o'
1682: 'w'
1683: 'n'
1684: 'T'
1685: 'y'
1686: 'p'
1687: 'e'
1688: 's'
1689: ':'
1690: '!'
1691: 'D'
1692: 'I'
1693: 'L'
1694: 'e'
1695: 'x'
1696: 'i'
1697: 'c'
1698: 'a'
1699: 'l'
1700: 'B'
1701: 'l'
1702: 'o'
1703: 'c'
1704: 'k'
1705: '!'
1706: 'D'
1707: 'I'
1708: 'L'
1709: 'e'
1710: 'x'
1711: 'i'
1712: 'c'
1713: 'a'
1714: 'l'
1715: 'B'
1716: 'l'
1717: 'o'
1718: 'c'
1719: 'k'
1720: 'F'
1721: 'i'
1722: 'l'
1723: 'e'
1724: '!'
1725: 'D'
1726: 'I'
1727: 'L'
1728: 'o'
1729: 'c'
1730: 'a'
1731: 't'
1732: 'i'
1733: 'o'
1734: 'n'
1735: 'i'
1736: 'n'
1737: 'l'
1738: 'i'
1739: 'n'
1740: 'e'
1741: 'd'
1742: 'A'
1743: 't'
1744: ':'
1745: '!'
1746: 'D'
1747: 'I'
1748: 'L'
1749: 'o'
1750: 'c'
1751: 'a'
1752: 'l'
1753: 'V'
1754: 'a'
1755: 'r'
1756: 'i'
1757: 'a'
1758: 'b'
1759: 'l'
1760: 'e'
1761: 'a'
1762: 'r'
1763: 'g'
1764: ':'
1765: '!'
1766: 'D'
1767: 'I'
1768: 'E'
1769: 'x'
1770: 'p'
1771: 'r'
1772: 'e'
1773: 's'
1774: 's'
1775: 'i'
1776: 'o'
1777: 'n'
1778: '!'
1779: 'D'
1780: 'I'
1781: 'G'
1782: 'l'
1783: 'o'
1784: 'b'
1785: 'a'
1786: 'l'
1787: 'V'
1788: 'a'
1789: 'r'
1790: 'i'
1791: 'a'
1792: 'b'
1793: 'l'
1794: 'e'
1795: 'E'
1796: 'x'
1797: 'p'
1798: 'r'
1799: 'e'
1800: 's'
1801: 's'
1802: 'i'
1803: 'o'
1804: 'n'
1805: 'v'
1806: 'a'
1807: 'r'
1808: ':'
1809: 'e'
1810: 'x'
1811: 'p'
1812: 'r'
1813: ':'
1814: '!'
1815: 'D'
1816: 'I'
1817: 'O'
1818: 'b'
1819: 'j'
1820: 'C'
1821: 'P'
1822: 'r'
1823: 'o'
1824: 'p'
1825: 'e'
1826: 'r'
1827: 't'
1828: 'y'
1829: 's'
1830: 'e'
1831: 't'
1832: 't'
1833: 'e'
1834: 'r'
1835: ':'
1836: 'g'
1837: 'e'
1838: 't'
1839: 't'
1840: 'e'
1841: 'r'
1842: ':'
1843: 'a'
1844: 't'
1845: 't'
1846: 'r'
1847: 'i'
1848: 'b'
1849: 'u'
1850: 't'
1851: 'e'
1852: 's'
1853: ':'
1854: '!'
1855: 'D'
1856: 'I'
1857: 'I'
1858: 'm'
1859: 'p'
1860: 'o'
1861: 'r'
1862: 't'
1863: 'e'
1864: 'd'
1865: 'E'
1866: 'n'
1867: 't'
1868: 'i'
1869: 't'
1870: 'y'
1871: 'e'
1872: 'n'
1873: 't'
1874: 'i'
1875: 't'
1876: 'y'
1877: ':'
1878: '!'
1879: 'D'
1880: 'I'
1881: 'M'
1882: 'a'
1883: 'c'
1884: 'r'
1885: 'o'
1886: '!'
1887: 'D'
1888: 'I'
1889: 'M'
1890: 'a'
1891: 'c'
1892: 'r'
1893: 'o'
1894: 'F'
1895: 'i'
1896: 'l'
1897: 'e'
1898: 'n'
1899: 'o'
1900: 'd'
1901: 'e'
1902: 's'
1903: ':'
1904: '!'
1905: 'G'
1906: 'e'
1907: 'n'
1908: 'e'
1909: 'r'
1910: 'i'
1911: 'c'
1912: 'D'
1913: 'I'
1914: 'N'
1915: 'o'
1916: 'd'
1917: 'e'
1918: 'h'
1919: 'e'
1920: 'a'
1921: 'd'
1922: 'e'
1923: 'r'
1924: ':'
1925: 'o'
1926: 'p'
1927: 'e'
1928: 'r'
1929: 'a'
1930: 'n'
1931: 'd'
1932: 's'
1933: ':'
1934: 'f'
1935: 'i'
1936: 'l'
1937: 'e'
1938: ':'
1939: 'i'
1940: 's'
1941: 'O'
1942: 'p'
1943: 't'
1944: 'i'
1945: 'm'
1946: 'i'
1947: 'z'
1948: 'e'
1949: 'd'
1950: ':'
1951: 't'
1952: 'a'
1953: 'g'
1954: ':'
1955: 'n'
1956: 'a'
1957: 'm'
1958: 'e'
1959: ':'
1960: 's'
1961: 'i'
1962: 'z'
1963: 'e'
1964: ':'
1965: 'a'
1966: 'l'
1967: 'i'
1968: 'g'
1969: 'n'
1970: ':'
1971: 'l'
1972: 'i'
1973: 'n'
1974: 'e'
1975: ':'
1976: 's'
1977: 'c'
1978: 'o'
1979: 'p'
1980: 'e'
1981: ':'
1982: 'b'
1983: 'a'
1984: 's'
1985: 'e'
1986: 'T'
1987: 'y'
1988: 'p'
1989: 'e'
1990: ':'
1991: 'o'
1992: 'f'
1993: 'f'
1994: 's'
1995: 'e'
1996: 't'
1997: ':'
1998: 't'
1999: 'e'
2000: 'm'
2001: 'p'
2002: 'l'
2003: 'a'
2004: 't'
2005: 'e'
2006: 'P'
2007: 'a'
2008: 'r'
2009: 'a'
2010: 'm'
2011: 's'
2012: ':'
2013: 't'
2014: 'y'
2015: 'p'
2016: 'e'
2017: ':'
2018: 'l'
2019: 'i'
2020: 'n'
2021: 'k'
2022: 'a'
2023: 'g'
2024: 'e'
2025: 'N'
2026: 'a'
2027: 'm'
2028: 'e'
2029: ':'
2030: 'i'
2031: 's'
2032: 'L'
2033: 'o'
2034: 'c'
2035: 'a'
2036: 'l'
2037: ':'
2038: 'i'
2039: 's'
2040: 'D'
2041: 'e'
2042: 'f'
2043: 'i'
2044: 'n'
2045: 'i'
2046: 't'
2047: 'i'
2048: 'o'
2049: 'n'
2050: ':'
2051: 'd'
2052: 'e'
2053: 'c'
2054: 'l'
2055: 'a'
2056: 'r'
2057: 'a'
2058: 't'
2059: 'i'
2060: 'o'
2061: 'n'
2062: ':'
2063: 'c'
2064: 'o'
2065: 'l'
2066: 'u'
2067: 'm'
2068: 'n'
2069: ':'
2070: '|'
2071: 'F'
2072: 'u'
2073: 'l'
2074: 'l'
2075: 'D'
2076: 'e'
2077: 'b'
2078: 'u'
2079: 'g'
2080: 'L'
2081: 'i'
2082: 'n'
2083: 'e'
2084: 'T'
2085: 'a'
2086: 'b'
2087: 'l'
2088: 'e'
2089: 's'
2090: 'O'
2091: 'n'
2092: 'l'
2093: 'y'
2094: 'N'
2095: 'o'
2096: 'D'
2097: 'e'
2098: 'b'
2099: 'u'
2100: 'g'
2101: 'a'
2102: 'l'
2103: 'i'
2104: 'g'
2105: 'n'
2106: 'a'
2107: 'l'
2108: 'l'
2109: 'o'
2110: 'c'
2111: 's'
2112: 'i'
2113: 'z'
2114: 'e'
2115: '.'
2116: '.'
2117: '.'
2118: 'a'
2119: 'c'
2120: 'q'
2121: '_'
2122: 'r'
2123: 'e'
2124: 'l'
2125: 'a'
2126: 'c'
2127: 'q'
2128: 'u'
2129: 'i'
2130: 'r'
2131: 'e'
2132: 'm'
2133: 'o'
2134: 'n'
2135: 'o'
2136: 't'
2137: 'o'
2138: 'n'
2139: 'i'
2140: 'c'
2141: 'r'
2142: 'e'
2143: 'l'
2144: 'e'
2145: 'a'
2146: 's'
2147: 'e'
2148: 's'
2149: 'e'
2150: 'q'
2151: '_'
2152: 'c'
2153: 's'
2154: 't'
2155: 'u'
2156: 'n'
2157: 'o'
2158: 'r'
2159: 'd'
2160: 'e'
2161: 'r'
2162: 'e'
2163: 'd'
2164: 'a'
2165: 'm'
2166: 'd'
2167: 'g'
2168: 'p'
2169: 'u'
2170: '_'
2171: 'c'
2172: 's'
2173: 'a'
2174: 'm'
2175: 'd'
2176: 'g'
2177: 'p'
2178: 'u'
2179: '_'
2180: 'e'
2181: 's'
2182: 'a'
2183: 'm'
2184: 'd'
2185: 'g'
2186: 'p'
2187: 'u'
2188: '_'
2189: 'g'
2190: 's'
2191: 'a'
2192: 'm'
2193: 'd'
2194: 'g'
2195: 'p'
2196: 'u'
2197: '_'
2198: 'h'
2199: 's'
2200: 'a'
2201: 'm'
2202: 'd'
2203: 'g'
2204: 'p'
2205: 'u'
2206: '_'
2207: 'k'
2208: 'e'
2209: 'r'
2210: 'n'
2211: 'e'
2212: 'l'
2213: 'a'
2214: 'm'
2215: 'd'
2216: 'g'
2217: 'p'
2218: 'u'
2219: '_'
2220: 'l'
2221: 's'
2222: 'a'
2223: 'm'
2224: 'd'
2225: 'g'
2226: 'p'
2227: 'u'
2228: '_'
2229: 'p'
2230: 's'
2231: 'a'
2232: 'm'
2233: 'd'
2234: 'g'
2235: 'p'
2236: 'u'
2237: '_'
2238: 'v'
2239: 's'
2240: 'a'
2241: 'n'
2242: 'y'
2243: 'r'
2244: 'e'
2245: 'g'
2246: 'c'
2247: 'c'
2248: 'a'
2249: 'r'
2250: 'm'
2251: '_'
2252: 'a'
2253: 'a'
2254: 'p'
2255: 'c'
2256: 's'
2257: '_'
2258: 'v'
2259: 'f'
2260: 'p'
2261: 'c'
2262: 'c'
2263: 'a'
2264: 'r'
2265: 'm'
2266: '_'
2267: 'a'
2268: 'a'
2269: 'p'
2270: 'c'
2271: 's'
2272: 'c'
2273: 'c'
2274: 'a'
2275: 'r'
2276: 'm'
2277: '_'
2278: 'a'
2279: 'p'
2280: 'c'
2281: 's'
2282: 'c'
2283: 'c'
2284: 'a'
2285: 'v'
2286: 'r'
2287: '_'
2288: 'i'
2289: 'n'
2290: 't'
2291: 'r'
2292: 'c'
2293: 'c'
2294: 'a'
2295: 'v'
2296: 'r'
2297: '_'
2298: 's'
2299: 'i'
2300: 'g'
2301: 'n'
2302: 'a'
2303: 'l'
2304: 'c'
2305: 'c'
2306: 'c'
2307: 'c'
2308: 'c'
2309: 'c'
2310: 'o'
2311: 'l'
2312: 'd'
2313: 'c'
2314: 'c'
2315: 'c'
2316: 'x'
2317: 'x'
2318: '_'
2319: 'f'
2320: 'a'
2321: 's'
2322: 't'
2323: '_'
2324: 't'
2325: 'l'
2326: 's'
2327: 'c'
2328: 'c'
2329: 'f'
2330: 'a'
2331: 's'
2332: 't'
2333: 'c'
2334: 'c'
2335: 'g'
2336: 'h'
2337: 'c'
2338: 'c'
2339: 'c'
2340: 'h'
2341: 'h'
2342: 'v'
2343: 'm'
2344: '_'
2345: 'c'
2346: 'c'
2347: 'c'
2348: 'h'
2349: 'h'
2350: 'v'
2351: 'm'
2352: 'c'
2353: 'c'
2354: 'i'
2355: 'n'
2356: 't'
2357: 'e'
2358: 'l'
2359: '_'
2360: 'o'
2361: 'c'
2362: 'l'
2363: '_'
2364: 'b'
2365: 'i'
2366: 'c'
2367: 'c'
2368: 'm'
2369: 's'
2370: 'p'
2371: '4'
2372: '3'
2373: '0'
2374: '_'
2375: 'i'
2376: 'n'
2377: 't'
2378: 'r'
2379: 'c'
2380: 'c'
2381: 'p'
2382: 'r'
2383: 'e'
2384: 's'
2385: 'e'
2386: 'r'
2387: 'v'
2388: 'e'
2389: '_'
2390: 'a'
2391: 'l'
2392: 'l'
2393: 'c'
2394: 'c'
2395: 'p'
2396: 'r'
2397: 'e'
2398: 's'
2399: 'e'
2400: 'r'
2401: 'v'
2402: 'e'
2403: '_'
2404: 'm'
2405: 'o'
2406: 's'
2407: 't'
2408: 'c'
2409: 'c'
2410: 'p'
2411: 't'
2412: 'x'
2413: '_'
2414: 'd'
2415: 'e'
2416: 'v'
2417: 'i'
2418: 'c'
2419: 'e'
2420: 'p'
2421: 't'
2422: 'x'
2423: '_'
2424: 'k'
2425: 'e'
2426: 'r'
2427: 'n'
2428: 'e'
2429: 'l'
2430: 's'
2431: 'p'
2432: 'i'
2433: 'r'
2434: '_'
2435: 'f'
2436: 'u'
2437: 'n'
2438: 'c'
2439: 's'
2440: 'p'
2441: 'i'
2442: 'r'
2443: '_'
2444: 'k'
2445: 'e'
2446: 'r'
2447: 'n'
2448: 'e'
2449: 'l'
2450: 's'
2451: 'w'
2452: 'i'
2453: 'f'
2454: 't'
2455: 'c'
2456: 'c'
2457: 'w'
2458: 'e'
2459: 'b'
2460: 'k'
2461: 'i'
2462: 't'
2463: '_'
2464: 'j'
2465: 's'
2466: 'c'
2467: 'c'
2468: 'w'
2469: 'i'
2470: 'n'
2471: '6'
2472: '4'
2473: 'c'
2474: 'c'
2475: 'x'
2476: '8'
2477: '6'
2478: '_'
2479: '6'
2480: '4'
2481: '_'
2482: 's'
2483: 'y'
2484: 's'
2485: 'v'
2486: 'c'
2487: 'c'
2488: 'x'
2489: '8'
2490: '6'
2491: '_'
2492: 'f'
2493: 'a'
2494: 's'
2495: 't'
2496: 'c'
2497: 'a'
2498: 'l'
2499: 'l'
2500: 'c'
2501: 'c'
2502: 'x'
2503: '8'
2504: '6'
2505: '_'
2506: 'i'
2507: 'n'
2508: 't'
2509: 'r'
2510: 'c'
2511: 'c'
2512: 'x'
2513: '8'
2514: '6'
2515: '_'
2516: 'r'
2517: 'e'
2518: 'g'
2519: 'c'
2520: 'a'
2521: 'l'
2522: 'l'
2523: 'c'
2524: 'c'
2525: 'x'
2526: '8'
2527: '6'
2528: '_'
2529: 's'
2530: 't'
2531: 'd'
2532: 'c'
2533: 'a'
2534: 'l'
2535: 'l'
2536: 'c'
2537: 'c'
2538: 'x'
2539: '8'
2540: '6'
2541: '_'
2542: 't'
2543: 'h'
2544: 'i'
2545: 's'
2546: 'c'
2547: 'a'
2548: 'l'
2549: 'l'
2550: 'c'
2551: 'c'
2552: 'x'
2553: '8'
2554: '6'
2555: '_'
2556: 'v'
2557: 'e'
2558: 'c'
2559: 't'
2560: 'o'
2561: 'r'
2562: 'c'
2563: 'a'
2564: 'l'
2565: 'l'
2566: 'c'
2567: 'c'
2568: 'c'
2569: 'c'
2570: 'd'
2571: 'e'
2572: 'r'
2573: 'e'
2574: 'f'
2575: 'e'
2576: 'r'
2577: 'e'
2578: 'n'
2579: 'c'
2580: 'e'
2581: 'a'
2582: 'b'
2583: 'l'
2584: 'e'
2585: 'd'
2586: 'e'
2587: 'r'
2588: 'e'
2589: 'f'
2590: 'e'
2591: 'r'
2592: 'e'
2593: 'n'
2594: 'c'
2595: 'e'
2596: 'a'
2597: 'b'
2598: 'l'
2599: 'e'
2600: '_'
2601: 'o'
2602: 'r'
2603: '_'
2604: 'n'
2605: 'u'
2606: 'l'
2607: 'l'
2608: 'd'
2609: 'l'
2610: 'l'
2611: 'e'
2612: 'x'
2613: 'p'
2614: 'o'
2615: 'r'
2616: 't'
2617: 'd'
2618: 'l'
2619: 'l'
2620: 'i'
2621: 'm'
2622: 'p'
2623: 'o'
2624: 'r'
2625: 't'
2626: 'e'
2627: 'x'
2628: 'a'
2629: 'c'
2630: 't'
2631: 'a'
2632: 'f'
2633: 'n'
2634: 'a'
2635: 'r'
2636: 'c'
2637: 'p'
2638: 'c'
2639: 'o'
2640: 'n'
2641: 't'
2642: 'r'
2643: 'a'
2644: 'c'
2645: 't'
2646: 'f'
2647: 'a'
2648: 's'
2649: 't'
2650: 'n'
2651: 'i'
2652: 'n'
2653: 'f'
2654: 'n'
2655: 'n'
2656: 'a'
2657: 'n'
2658: 'n'
2659: 's'
2660: 'z'
2661: 'r'
2662: 'e'
2663: 'a'
2664: 's'
2665: 's'
2666: 'o'
2667: 'c'
2668: 'o'
2669: 'e'
2670: 'q'
2671: 'o'
2672: 'g'
2673: 'e'
2674: 'o'
2675: 'g'
2676: 't'
2677: 'o'
2678: 'l'
2679: 'e'
2680: 'o'
2681: 'l'
2682: 't'
2683: 'o'
2684: 'n'
2685: 'e'
2686: 'o'
2687: 'r'
2688: 'd'
2689: 'u'
2690: 'e'
2691: 'q'
2692: 'u'
2693: 'g'
2694: 'e'
2695: 'u'
2696: 'g'
2697: 't'
2698: 'u'
2699: 'l'
2700: 'e'
2701: 'u'
2702: 'l'
2703: 't'
2704: 'u'
2705: 'n'
2706: 'e'
2707: 'u'
2708: 'n'
2709: 'o'
2710: 'a'
2711: 'l'
2712: 'w'
2713: 'a'
2714: 'y'
2715: 's'
2716: 'i'
2717: 'n'
2718: 'l'
2719: 'i'
2720: 'n'
2721: 'e'
2722: 'a'
2723: 'r'
2724: 'g'
2725: 'm'
2726: 'e'
2727: 'm'
2728: 'o'
2729: 'n'
2730: 'l'
2731: 'y'
2732: 'b'
2733: 'u'
2734: 'i'
2735: 'l'
2736: 't'
2737: 'i'
2738: 'n'
2739: 'c'
2740: 'o'
2741: 'l'
2742: 'd'
2743: 'c'
2744: 'o'
2745: 'n'
2746: 'v'
2747: 'e'
2748: 'r'
2749: 'g'
2750: 'e'
2751: 'n'
2752: 't'
2753: 'i'
2754: 'n'
2755: 'a'
2756: 'c'
2757: 'c'
2758: 'e'
2759: 's'
2760: 's'
2761: 'i'
2762: 'b'
2763: 'l'
2764: 'e'
2765: 'm'
2766: 'e'
2767: 'm'
2768: '_'
2769: 'o'
2770: 'r'
2771: '_'
2772: 'a'
2773: 'r'
2774: 'g'
2775: 'm'
2776: 'e'
2777: 'm'
2778: 'o'
2779: 'n'
2780: 'l'
2781: 'y'
2782: 'i'
2783: 'n'
2784: 'a'
2785: 'c'
2786: 'c'
2787: 'e'
2788: 's'
2789: 's'
2790: 'i'
2791: 'b'
2792: 'l'
2793: 'e'
2794: 'm'
2795: 'e'
2796: 'm'
2797: 'o'
2798: 'n'
2799: 'l'
2800: 'y'
2801: 'i'
2802: 'n'
2803: 'l'
2804: 'i'
2805: 'n'
2806: 'e'
2807: 'h'
2808: 'i'
2809: 'n'
2810: 't'
2811: 'j'
2812: 'u'
2813: 'm'
2814: 'p'
2815: 't'
2816: 'a'
2817: 'b'
2818: 'l'
2819: 'e'
2820: 'm'
2821: 'i'
2822: 'n'
2823: 's'
2824: 'i'
2825: 'z'
2826: 'e'
2827: 'n'
2828: 'a'
2829: 'k'
2830: 'e'
2831: 'd'
2832: 'n'
2833: 'o'
2834: 'b'
2835: 'u'
2836: 'i'
2837: 'l'
2838: 't'
2839: 'i'
2840: 'n'
2841: 'n'
2842: 'o'
2843: 'd'
2844: 'u'
2845: 'p'
2846: 'l'
2847: 'i'
2848: 'c'
2849: 'a'
2850: 't'
2851: 'e'
2852: 'n'
2853: 'o'
2854: 'i'
2855: 'm'
2856: 'p'
2857: 'l'
2858: 'i'
2859: 'c'
2860: 'i'
2861: 't'
2862: 'f'
2863: 'l'
2864: 'o'
2865: 'a'
2866: 't'
2867: 'n'
2868: 'o'
2869: 'i'
2870: 'n'
2871: 'l'
2872: 'i'
2873: 'n'
2874: 'e'
2875: 'n'
2876: 'o'
2877: 'n'
2878: 'l'
2879: 'a'
2880: 'z'
2881: 'y'
2882: 'b'
2883: 'i'
2884: 'n'
2885: 'd'
2886: 'n'
2887: 'o'
2888: 'r'
2889: 'e'
2890: 'c'
2891: 'u'
2892: 'r'
2893: 's'
2894: 'e'
2895: 'n'
2896: 'o'
2897: 'r'
2898: 'e'
2899: 'd'
2900: 'z'
2901: 'o'
2902: 'n'
2903: 'e'
2904: 'n'
2905: 'o'
2906: 'r'
2907: 'e'
2908: 't'
2909: 'u'
2910: 'r'
2911: 'n'
2912: 'n'
2913: 'o'
2914: 'u'
2915: 'n'
2916: 'w'
2917: 'i'
2918: 'n'
2919: 'd'
2920: 'o'
2921: 'p'
2922: 't'
2923: 'n'
2924: 'o'
2925: 'n'
2926: 'e'
2927: 'o'
2928: 'p'
2929: 't'
2930: 's'
2931: 'i'
2932: 'z'
2933: 'e'
2934: 'r'
2935: 'e'
2936: 'a'
2937: 'd'
2938: 'n'
2939: 'o'
2940: 'n'
2941: 'e'
2942: 'r'
2943: 'e'
2944: 'a'
2945: 'd'
2946: 'o'
2947: 'n'
2948: 'l'
2949: 'y'
2950: 'r'
2951: 'e'
2952: 't'
2953: 'u'
2954: 'r'
2955: 'n'
2956: 's'
2957: '_'
2958: 't'
2959: 'w'
2960: 'i'
2961: 'c'
2962: 'e'
2963: 's'
2964: 'a'
2965: 'f'
2966: 'e'
2967: 's'
2968: 't'
2969: 'a'
2970: 'c'
2971: 'k'
2972: 's'
2973: 'a'
2974: 'n'
2975: 'i'
2976: 't'
2977: 'i'
2978: 'z'
2979: 'e'
2980: '_'
2981: 'a'
2982: 'd'
2983: 'd'
2984: 'r'
2985: 'e'
2986: 's'
2987: 's'
2988: 's'
2989: 'a'
2990: 'n'
2991: 'i'
2992: 't'
2993: 'i'
2994: 'z'
2995: 'e'
2996: '_'
2997: 'h'
2998: 'w'
2999: 'a'
3000: 'd'
3001: 'd'
3002: 'r'
3003: 'e'
3004: 's'
3005: 's'
3006: 's'
3007: 'a'
3008: 'n'
3009: 'i'
3010: 't'
3011: 'i'
3012: 'z'
3013: 'e'
3014: '_'
3015: 'm'
3016: 'e'
3017: 'm'
3018: 'o'
3019: 'r'
3020: 'y'
3021: 's'
3022: 'a'
3023: 'n'
3024: 'i'
3025: 't'
3026: 'i'
3027: 'z'
3028: 'e'
3029: '_'
3030: 't'
3031: 'h'
3032: 'r'
3033: 'e'
3034: 'a'
3035: 'd'
3036: 's'
3037: 'p'
3038: 'e'
3039: 'c'
3040: 'u'
3041: 'l'
3042: 'a'
3043: 't'
3044: 'a'
3045: 'b'
3046: 'l'
3047: 'e'
3048: 's'
3049: 's'
3050: 'p'
3051: 's'
3052: 's'
3053: 'p'
3054: 'r'
3055: 'e'
3056: 'q'
3057: 's'
3058: 's'
3059: 'p'
3060: 's'
3061: 't'
3062: 'r'
3063: 'o'
3064: 'n'
3065: 'g'
3066: 's'
3067: 't'
3068: 'r'
3069: 'i'
3070: 'c'
3071: 't'
3072: 'f'
3073: 'p'
3074: 'u'
3075: 'w'
3076: 't'
3077: 'a'
3078: 'b'
3079: 'l'
3080: 'e'
3081: 'w'
3082: 'r'
3083: 'i'
3084: 't'
3085: 'e'
3086: 'o'
3087: 'n'
3088: 'l'
3089: 'y'
3090: 'i'
3091: 'n'
3092: 'b'
3093: 'o'
3094: 'u'
3095: 'n'
3096: 'd'
3097: 's'
3098: 'e'
3099: 'q'
3100: 'n'
3101: 'e'
3102: 's'
3103: 'g'
3104: 'e'
3105: 's'
3106: 'g'
3107: 't'
3108: 's'
3109: 'l'
3110: 'e'
3111: 's'
3112: 'l'
3113: 't'
3114: 'a'
3115: 'p'
3116: 'p'
3117: 'e'
3118: 'n'
3119: 'd'
3120: 'i'
3121: 'n'
3122: 'g'
3123: 'a'
3124: 'v'
3125: 'a'
3126: 'i'
3127: 'l'
3128: 'a'
3129: 'b'
3130: 'l'
3131: 'e'
3132: '_'
3133: 'e'
3134: 'x'
3135: 't'
3136: 'e'
3137: 'r'
3138: 'n'
3139: 'a'
3140: 'l'
3141: 'l'
3142: 'y'
3143: 'c'
3144: 'o'
3145: 'm'
3146: 'm'
3147: 'o'
3148: 'n'
3149: 'i'
3150: 'n'
3151: 't'
3152: 'e'
3153: 'r'
3154: 'n'
3155: 'a'
3156: 'l'
3157: 'l'
3158: 'i'
3159: 'n'
3160: 'k'
3161: 'o'
3162: 'n'
3163: 'c'
3164: 'e'
3165: 'l'
3166: 'i'
3167: 'n'
3168: 'k'
3169: 'o'
3170: 'n'
3171: 'c'
3172: 'e'
3173: '_'
3174: 'o'
3175: 'd'
3176: 'r'
3177: 'p'
3178: 'r'
3179: 'i'
3180: 'v'
3181: 'a'
3182: 't'
3183: 'e'
3184: 'w'
3185: 'e'
3186: 'a'
3187: 'k'
3188: '_'
3189: 'o'
3190: 'd'
3191: 'r'
3192: 'e'
3193: 'x'
3194: 't'
3195: 'e'
3196: 'r'
3197: 'n'
3198: '_'
3199: 'w'
3200: 'e'
3201: 'a'
3202: 'k'
3203: 'e'
3204: 'x'
3205: 't'
3206: 'e'
3207: 'r'
3208: 'n'
3209: 'a'
3210: 'l'
3211: 'n'
3212: 's'
3213: 'w'
3214: 'n'
3215: 'u'
3216: 'w'
3217: 'b'
3218: 'y'
3219: 'v'
3220: 'a'
3221: 'l'
3222: 'i'
3223: 'n'
3224: 'r'
3225: 'e'
3226: 'g'
3227: 'n'
3228: 'e'
3229: 's'
3230: 't'
3231: 'n'
3232: 'o'
3233: 'a'
3234: 'l'
3235: 'i'
3236: 'a'
3237: 's'
3238: 'n'
3239: 'o'
3240: 'c'
3241: 'a'
3242: 'p'
3243: 't'
3244: 'u'
3245: 'r'
3246: 'e'
3247: 'n'
3248: 'o'
3249: 'n'
3250: 'n'
3251: 'u'
3252: 'l'
3253: 'l'
3254: 'r'
3255: 'e'
3256: 't'
3257: 'u'
3258: 'r'
3259: 'n'
3260: 'e'
3261: 'd'
3262: 's'
3263: 'i'
3264: 'g'
3265: 'n'
3266: 'e'
3267: 'x'
3268: 't'
3269: 's'
3270: 'r'
3271: 'e'
3272: 't'
3273: 's'
3274: 'w'
3275: 'i'
3276: 'f'
3277: 't'
3278: 's'
3279: 'e'
3280: 'l'
3281: 'f'
3282: 'z'
3283: 'e'
3284: 'r'
3285: 'o'
3286: 'e'
3287: 'x'
3288: 't'
3289: 'd'
3290: 's'
3291: 'o'
3292: '_'
3293: 'l'
3294: 'o'
3295: 'c'
3296: 'a'
3297: 'l'
3298: 'd'
3299: 's'
3300: 'o'
3301: '_'
3302: 'p'
3303: 'r'
3304: 'e'
3305: 'e'
3306: 'm'
3307: 'p'
3308: 't'
3309: 'a'
3310: 'b'
3311: 'l'
3312: 'e'
3313: 's'
3314: 'e'
3315: 'c'
3316: 't'
3317: 'i'
3318: 'o'
3319: 'n'
3320: 's'
3321: 'y'
3322: 'n'
3323: 'c'
3324: 's'
3325: 'c'
3326: 'o'
3327: 'p'
3328: 'e'
3329: 't'
3330: 'h'
3331: 'r'
3332: 'e'
3333: 'a'
3334: 'd'
3335: '_'
3336: 'l'
3337: 'o'
3338: 'c'
3339: 'a'
3340: 'l'
3341: 'i'
3342: 'n'
3343: 'i'
3344: 't'
3345: 'i'
3346: 'a'
3347: 'l'
3348: 'e'
3349: 'x'
3350: 'e'
3351: 'c'
3352: 'l'
3353: 'o'
3354: 'c'
3355: 'a'
3356: 'l'
3357: 'd'
3358: 'y'
3359: 'n'
3360: 'a'
3361: 'm'
3362: 'i'
3363: 'c'
3364: 'l'
3365: 'o'
3366: 'c'
3367: 'a'
3368: 'l'
3369: 'e'
3370: 'x'
3371: 'e'
3372: 'c'
3373: 'l'
3374: 'o'
3375: 'c'
3376: 'a'
3377: 'l'
3378: '_'
3379: 'u'
3380: 'n'
3381: 'n'
3382: 'a'
3383: 'm'
3384: 'e'
3385: 'd'
3386: '_'
3387: 'a'
3388: 'd'
3389: 'd'
3390: 'r'
3391: 'u'
3392: 'n'
3393: 'n'
3394: 'a'
3395: 'm'
3396: 'e'
3397: 'd'
3398: '_'
3399: 'a'
3400: 'd'
3401: 'd'
3402: 'r'
3403: 'd'
3404: 'e'
3405: 'f'
3406: 'a'
3407: 'u'
3408: 'l'
3409: 't'
3410: 'h'
3411: 'i'
3412: 'd'
3413: 'd'
3414: 'e'
3415: 'n'
3416: 'p'
3417: 'r'
3418: 'o'
3419: 't'
3420: 'e'
3421: 'c'
3422: 't'
3423: 'e'
3424: 'd'
3425: 'v'
3426: 'o'
3427: 'l'
3428: 'a'
3429: 't'
3430: 'i'
3431: 'l'
3432: 'e'
3433: '$'
3434: '-'
3435: '.'
3436: '_'
3437: '\'
3438: '@'
3439: '@'
3440: '%'
3441: '%'
3442: '-'
3443: '.'
3444: '+'
3445: '-'
3446: 'e'
3447: 'E'
3448: '0'
3449: 'x'
3450: '0'
3451: 'x'
3452: 'K'
3453: '0'
3454: 'x'
3455: 'L'
3456: '0'
3457: 'x'
3458: 'M'
3459: '0'
3460: 'x'
3461: 'H'
3462: '"'
3463: '"'
3464: ';'
3465: '\n'
3466: \u0000
3467: ' '
3468: '\t'
3469: '\r'
3470: '\n'
3471: 'A'-'Z'
3472: 'a'-'z'
3473: '0'-'9'
3474: 'A'-'F'
3475: 'a'-'f'
3476: .
*/
