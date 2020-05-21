package pic

// type yySymType struct{
// 	yys int
// 	str string
// 	result Result
// }

// /*
//  * Changes by Gunnar Ritter, Freiburg i. Br., Germany, October 2005.
//  *
//  * Derived from Plan 9 v4 /sys/src/cmd/pic/
//  *
//  * Copyright (C) 2003, Lucent Technologies Inc. and others.
//  * All Rights Reserved.
//  *
//  * Distributed under the terms of the Lucent Public License Version 1.02.
//  */

// /*	Sccsid @(#)picy.y	1.4 (gritter) 11/28/05	*/

// // #include <stdio.h>
// // #include "pic.h"
// // #include <math.h>
// // #include <stdlib.h>
// // #include <string.h>

// // #ifndef	RAND_MAX
// // #define	RAND_MAX	32767
// // #endif

// // YYSTYPE	y;

// // extern	void	yyerror(char *);
// // extern	int	yylex(void);
// const BOX = 1
// const LINE = 2
// const ARROW = 3
// const CIRCLE = 4
// const ELLIPSE = 5
// const ARC = 6
// const SPLINE = 7
// const BLOCK = 8
// const TEXT = 9
// const TROFF = 10
// const MOVE = 11
// const BLOCKEND = 12
// const PLACE = 13
// const PRINT = 57359
// const RESET = 57360
// const THRU = 57361
// const UNTIL = 57362
// const FOR = 57363
// const IF = 57364
// const COPY = 57365
// const THENSTR = 57366
// const ELSESTR = 57367
// const DOSTR = 57368
// const PLACENAME = 57369
// const VARNAME = 57370
// const SPRINTF = 57371
// const DEFNAME = 57372
// const ATTR = 57373
// const TEXTATTR = 57374
// const LEFT = 57375
// const RIGHT = 57376
// const UP = 57377
// const DOWN = 57378
// const FROM = 57379
// const TO = 57380
// const AT = 57381
// const BY = 57382
// const WITH = 57383
// const HEAD = 57384
// const CW = 57385
// const CCW = 57386
// const THEN = 57387
// const HEIGHT = 57388
// const WIDTH = 57389
// const RADIUS = 57390
// const DIAMETER = 57391
// const LENGTH = 57392
// const SIZE = 57393
// const CORNER = 57394
// const HERE = 57395
// const LAST = 57396
// const NTH = 57397
// const SAME = 57398
// const BETWEEN = 57399
// const AND = 57400
// const EAST = 57401
// const WEST = 57402
// const NORTH = 57403
// const SOUTH = 57404
// const NE = 57405
// const NW = 57406
// const SE = 57407
// const SW = 57408
// const START = 57409
// const END = 57410
// const DOTX = 57411
// const DOTY = 57412
// const DOTHT = 57413
// const DOTWID = 57414
// const DOTRAD = 57415
// const NUMBER = 57416
// const LOG = 57417
// const EXP = 57418
// const SIN = 57419
// const COS = 57420
// const ATAN2 = 57421
// const SQRT = 57422
// const RAND = 57423
// const MIN = 57424
// const MAX = 57425
// const INT = 57426
// const DIR = 57427
// const DOT = 57428
// const DASH = 57429
// const CHOP = 57430
// const FILL = 57431
// const NOEDGE = 57432
// const ST = 57433
// const OROR = 57434
// const ANDAND = 57435
// const GT = 57436
// const LT = 57437
// const LE = 57438
// const GE = 57439
// const EQ = 57440
// const NEQ = 57441
// const UMINUS = 57442
// const NOT = 57443

// var yyToknames = [...]string{
// 	"$end",
// 	"error",
// 	"$unk",
// 	"BOX",
// 	"LINE",
// 	"ARROW",
// 	"CIRCLE",
// 	"ELLIPSE",
// 	"ARC",
// 	"SPLINE",
// 	"BLOCK",
// 	"TEXT",
// 	"TROFF",
// 	"MOVE",
// 	"BLOCKEND",
// 	"PLACE",
// 	"PRINT",
// 	"RESET",
// 	"THRU",
// 	"UNTIL",
// 	"FOR",
// 	"IF",
// 	"COPY",
// 	"THENSTR",
// 	"ELSESTR",
// 	"DOSTR",
// 	"PLACENAME",
// 	"VARNAME",
// 	"SPRINTF",
// 	"DEFNAME",
// 	"ATTR",
// 	"TEXTATTR",
// 	"LEFT",
// 	"RIGHT",
// 	"UP",
// 	"DOWN",
// 	"FROM",
// 	"TO",
// 	"AT",
// 	"BY",
// 	"WITH",
// 	"HEAD",
// 	"CW",
// 	"CCW",
// 	"THEN",
// 	"HEIGHT",
// 	"WIDTH",
// 	"RADIUS",
// 	"DIAMETER",
// 	"LENGTH",
// 	"SIZE",
// 	"CORNER",
// 	"HERE",
// 	"LAST",
// 	"NTH",
// 	"SAME",
// 	"BETWEEN",
// 	"AND",
// 	"EAST",
// 	"WEST",
// 	"NORTH",
// 	"SOUTH",
// 	"NE",
// 	"NW",
// 	"SE",
// 	"SW",
// 	"START",
// 	"END",
// 	"DOTX",
// 	"DOTY",
// 	"DOTHT",
// 	"DOTWID",
// 	"DOTRAD",
// 	"NUMBER",
// 	"LOG",
// 	"EXP",
// 	"SIN",
// 	"COS",
// 	"ATAN2",
// 	"SQRT",
// 	"RAND",
// 	"MIN",
// 	"MAX",
// 	"INT",
// 	"DIR",
// 	"DOT",
// 	"DASH",
// 	"CHOP",
// 	"FILL",
// 	"NOEDGE",
// 	"ST",
// 	"'='",
// 	"OROR",
// 	"ANDAND",
// 	"GT",
// 	"LT",
// 	"LE",
// 	"GE",
// 	"EQ",
// 	"NEQ",
// 	"'+'",
// 	"'-'",
// 	"'*'",
// 	"'/'",
// 	"'%'",
// 	"UMINUS",
// 	"NOT",
// 	"'^'",
// 	"'}'",
// 	"':'",
// 	"','",
// 	"'{'",
// 	"']'",
// 	"'['",
// 	"'.'",
// 	"'('",
// 	"')'",
// }
// var yyStatenames = [...]string{}

// const yyEofCode = 1
// const yyErrCode = 2
// const yyInitialStackSize = 16

// var yyExca = [...]int{
// 	-1, 0,
// 	1, 2,
// 	-2, 0,
// 	-1, 1,
// 	1, -1,
// 	-2, 0,
// 	-1, 206,
// 	95, 0,
// 	96, 0,
// 	97, 0,
// 	98, 0,
// 	99, 0,
// 	100, 0,
// 	-2, 159,
// 	-1, 213,
// 	95, 0,
// 	96, 0,
// 	97, 0,
// 	98, 0,
// 	99, 0,
// 	100, 0,
// 	-2, 158,
// 	-1, 214,
// 	95, 0,
// 	96, 0,
// 	97, 0,
// 	98, 0,
// 	99, 0,
// 	100, 0,
// 	-2, 160,
// 	-1, 215,
// 	95, 0,
// 	96, 0,
// 	97, 0,
// 	98, 0,
// 	99, 0,
// 	100, 0,
// 	-2, 161,
// 	-1, 216,
// 	95, 0,
// 	96, 0,
// 	97, 0,
// 	98, 0,
// 	99, 0,
// 	100, 0,
// 	-2, 162,
// 	-1, 217,
// 	95, 0,
// 	96, 0,
// 	97, 0,
// 	98, 0,
// 	99, 0,
// 	100, 0,
// 	-2, 163,
// 	-1, 264,
// 	69, 112,
// 	70, 112,
// 	71, 112,
// 	72, 112,
// 	73, 112,
// 	-2, 85,
// 	-1, 270,
// 	95, 0,
// 	96, 0,
// 	97, 0,
// 	98, 0,
// 	99, 0,
// 	100, 0,
// 	-2, 159,
// }

// const yyPrivate = 57344

// const yyLast = 1654

// var yyAct = [...]int{

// 	173, 334, 139, 32, 53, 68, 312, 242, 124, 125,
// 	137, 42, 115, 197, 116, 117, 118, 119, 110, 111,
// 	112, 113, 114, 95, 227, 122, 162, 320, 81, 43,
// 	274, 137, 92, 319, 51, 299, 273, 161, 137, 160,
// 	159, 105, 158, 157, 156, 155, 154, 153, 98, 127,
// 	128, 131, 109, 298, 246, 235, 152, 149, 233, 40,
// 	112, 113, 114, 122, 51, 122, 41, 72, 40, 194,
// 	102, 164, 166, 138, 130, 110, 111, 112, 113, 114,
// 	129, 83, 122, 169, 190, 73, 74, 75, 76, 77,
// 	78, 79, 80, 276, 246, 200, 110, 111, 112, 113,
// 	114, 138, 126, 122, 124, 125, 107, 38, 204, 206,
// 	105, 208, 209, 210, 211, 212, 213, 214, 215, 216,
// 	217, 218, 219, 220, 195, 221, 224, 132, 133, 134,
// 	135, 136, 51, 51, 124, 125, 124, 125, 205, 207,
// 	198, 199, 34, 316, 275, 321, 168, 85, 223, 226,
// 	234, 124, 125, 44, 236, 237, 238, 239, 240, 241,
// 	203, 243, 244, 245, 232, 167, 170, 247, 249, 228,
// 	124, 125, 86, 252, 93, 253, 105, 105, 105, 105,
// 	105, 96, 97, 123, 261, 262, 263, 265, 335, 336,
// 	337, 338, 81, 124, 125, 267, 268, 192, 270, 51,
// 	51, 51, 51, 51, 251, 254, 255, 256, 257, 260,
// 	121, 120, 115, 197, 116, 117, 118, 119, 110, 111,
// 	112, 113, 114, 278, 90, 122, 280, 86, 311, 282,
// 	287, 193, 191, 283, 229, 132, 133, 134, 135, 136,
// 	71, 201, 142, 146, 147, 143, 144, 145, 148, 250,
// 	285, 35, 281, 300, 66, 67, 69, 284, 87, 88,
// 	287, 288, 35, 164, 166, 269, 2, 4, 36, 230,
// 	37, 285, 286, 39, 163, 305, 105, 105, 308, 36,
// 	310, 196, 188, 24, 171, 24, 266, 149, 84, 24,
// 	230, 231, 151, 82, 313, 70, 314, 315, 1, 51,
// 	51, 69, 165, 317, 318, 306, 307, 37, 100, 24,
// 	322, 5, 323, 142, 146, 147, 143, 144, 145, 148,
// 	248, 331, 24, 24, 26, 6, 12, 13, 14, 304,
// 	89, 339, 91, 0, 301, 340, 0, 0, 0, 0,
// 	341, 271, 272, 16, 20, 21, 17, 18, 19, 22,
// 	37, 35, 25, 23, 52, 46, 10, 11, 0, 0,
// 	30, 31, 29, 141, 0, 24, 103, 46, 36, 202,
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 66,
// 	67, 69, 54, 0, 24, 0, 0, 0, 0, 0,
// 	0, 66, 67, 69, 54, 0, 0, 0, 0, 0,
// 	0, 45, 56, 57, 58, 59, 60, 61, 62, 64,
// 	63, 65, 0, 45, 56, 57, 58, 59, 60, 61,
// 	62, 64, 63, 65, 9, 0, 0, 0, 49, 48,
// 	101, 0, 303, 0, 55, 0, 0, 0, 0, 0,
// 	49, 48, 35, 94, 0, 0, 55, 0, 0, 0,
// 	0, 27, 0, 33, 0, 50, 0, 52, 46, 36,
// 	0, 172, 181, 0, 0, 0, 0, 175, 176, 177,
// 	178, 179, 182, 0, 142, 146, 147, 143, 144, 145,
// 	148, 150, 66, 67, 69, 54, 180, 121, 120, 115,
// 	197, 116, 117, 118, 119, 110, 111, 112, 113, 114,
// 	0, 0, 122, 0, 45, 56, 57, 58, 59, 60,
// 	61, 62, 64, 63, 65, 174, 183, 184, 185, 186,
// 	187, 0, 0, 35, 151, 0, 0, 0, 0, 0,
// 	0, 49, 48, 0, 35, 0, 0, 55, 52, 46,
// 	36, 0, 0, 0, 0, 0, 94, 0, 0, 52,
// 	46, 36, 0, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 0, 66, 67, 69, 54, 0, 0, 0,
// 	0, 0, 0, 0, 66, 67, 69, 54, 0, 0,
// 	0, 0, 0, 0, 0, 45, 56, 57, 58, 59,
// 	60, 61, 62, 64, 63, 65, 45, 56, 57, 58,
// 	59, 60, 61, 62, 64, 63, 65, 52, 46, 0,
// 	0, 0, 49, 48, 0, 0, 0, 0, 55, 52,
// 	46, 0, 0, 49, 48, 0, 0, 94, 0, 55,
// 	0, 0, 258, 67, 69, 54, 0, 0, 50, 0,
// 	0, 0, 0, 0, 66, 67, 69, 54, 0, 0,
// 	0, 0, 0, 0, 45, 56, 57, 58, 59, 60,
// 	61, 62, 64, 63, 65, 0, 45, 56, 57, 58,
// 	59, 60, 61, 62, 64, 63, 65, 264, 46, 0,
// 	0, 49, 48, 0, 0, 0, 0, 55, 52, 46,
// 	0, 0, 0, 49, 48, 259, 50, 0, 0, 55,
// 	0, 0, 66, 67, 69, 54, 0, 0, 50, 0,
// 	0, 0, 0, 66, 67, 69, 54, 0, 0, 0,
// 	0, 0, 0, 0, 45, 56, 57, 58, 59, 60,
// 	61, 62, 64, 63, 65, 45, 56, 57, 58, 59,
// 	60, 61, 62, 64, 63, 65, 52, 46, 0, 0,
// 	0, 49, 48, 0, 0, 0, 0, 55, 0, 0,
// 	0, 0, 49, 48, 0, 0, 94, 0, 55, 0,
// 	0, 66, 67, 69, 54, 0, 0, 225, 120, 115,
// 	197, 116, 117, 118, 119, 110, 111, 112, 113, 114,
// 	0, 0, 122, 45, 56, 57, 58, 59, 60, 61,
// 	62, 64, 63, 65, 0, 16, 20, 21, 17, 18,
// 	19, 22, 0, 35, 25, 23, 0, 0, 10, 11,
// 	49, 48, 30, 31, 29, 0, 55, 0, 7, 28,
// 	36, 0, 0, 0, 0, 222, 16, 20, 21, 17,
// 	18, 19, 22, 0, 35, 25, 23, 0, 0, 10,
// 	11, 0, 0, 30, 31, 29, 0, 0, 0, 7,
// 	28, 36, 3, 0, 16, 20, 21, 17, 18, 19,
// 	22, 0, 35, 25, 23, 0, 0, 10, 11, 0,
// 	0, 30, 31, 29, 0, 0, 9, 7, 28, 36,
// 	0, 0, 15, 0, 0, 121, 120, 115, 197, 116,
// 	117, 118, 119, 110, 111, 112, 113, 114, 0, 0,
// 	122, 0, 0, 27, 189, 33, 0, 9, 0, 333,
// 	0, 0, 0, 15, 121, 120, 115, 197, 116, 117,
// 	118, 119, 110, 111, 112, 113, 114, 0, 0, 122,
// 	0, 99, 109, 0, 27, 9, 33, 0, 332, 0,
// 	0, 15, 16, 20, 21, 17, 18, 19, 22, 0,
// 	35, 25, 23, 0, 0, 10, 11, 0, 0, 30,
// 	31, 29, 27, 0, 33, 7, 28, 36, 121, 120,
// 	115, 108, 116, 117, 118, 119, 110, 111, 112, 113,
// 	114, 0, 0, 122, 0, 0, 107, 0, 0, 0,
// 	0, 0, 229, 121, 120, 115, 197, 116, 117, 118,
// 	119, 110, 111, 112, 113, 114, 0, 0, 122, 0,
// 	0, 309, 0, 0, 0, 0, 0, 229, 0, 0,
// 	0, 0, 0, 9, 0, 0, 0, 0, 0, 15,
// 	121, 120, 115, 197, 116, 117, 118, 119, 110, 111,
// 	112, 113, 114, 0, 0, 122, 0, 0, 0, 0,
// 	27, 0, 33, 0, 326, 121, 120, 115, 197, 116,
// 	117, 118, 119, 110, 111, 112, 113, 114, 0, 0,
// 	122, 142, 146, 147, 143, 144, 145, 148, 140, 325,
// 	121, 120, 115, 197, 116, 117, 118, 119, 110, 111,
// 	112, 113, 114, 0, 0, 122, 0, 0, 0, 0,
// 	0, 0, 0, 0, 324, 121, 120, 115, 197, 116,
// 	117, 118, 119, 110, 111, 112, 113, 114, 0, 0,
// 	122, 141, 0, 0, 0, 0, 0, 0, 0, 297,
// 	121, 120, 115, 197, 116, 117, 118, 119, 110, 111,
// 	112, 113, 114, 0, 0, 122, 0, 0, 0, 0,
// 	0, 0, 0, 0, 294, 121, 120, 115, 197, 116,
// 	117, 118, 119, 110, 111, 112, 113, 114, 0, 0,
// 	122, 0, 0, 0, 0, 0, 0, 0, 0, 292,
// 	121, 120, 115, 197, 116, 117, 118, 119, 110, 111,
// 	112, 113, 114, 0, 0, 122, 0, 0, 0, 0,
// 	0, 0, 0, 0, 291, 121, 120, 115, 197, 116,
// 	117, 118, 119, 110, 111, 112, 113, 114, 0, 0,
// 	122, 0, 0, 0, 0, 0, 0, 0, 0, 290,
// 	121, 120, 115, 197, 116, 117, 118, 119, 110, 111,
// 	112, 113, 114, 0, 109, 122, 0, 0, 0, 0,
// 	0, 0, 0, 0, 289, 121, 120, 115, 197, 116,
// 	117, 118, 119, 110, 111, 112, 113, 114, 109, 0,
// 	122, 0, 0, 0, 0, 0, 0, 0, 106, 229,
// 	121, 120, 115, 108, 116, 117, 118, 119, 110, 111,
// 	112, 113, 114, 0, 0, 122, 0, 0, 107, 0,
// 	0, 0, 0, 0, 121, 120, 115, 108, 116, 117,
// 	118, 119, 110, 111, 112, 113, 114, 0, 0, 122,
// 	0, 0, 107, 121, 120, 115, 197, 116, 117, 118,
// 	119, 110, 111, 112, 113, 114, 0, 0, 122, 0,
// 	0, 296, 121, 120, 115, 197, 116, 117, 118, 119,
// 	110, 111, 112, 113, 114, 0, 0, 122, 0, 0,
// 	295, 121, 120, 115, 197, 116, 117, 118, 119, 110,
// 	111, 112, 113, 114, 0, 0, 122, 343, 0, 293,
// 	121, 120, 115, 197, 116, 117, 118, 119, 110, 111,
// 	112, 113, 114, 342, 0, 122, 0, 0, 279, 121,
// 	120, 115, 197, 116, 117, 118, 119, 110, 111, 112,
// 	113, 114, 330, 0, 122, 0, 0, 277, 0, 0,
// 	0, 0, 0, 0, 0, 0, 329, 0, 328, 0,
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 327, 0, 121, 120, 115, 197, 116, 117,
// 	118, 119, 110, 111, 112, 113, 114, 302, 0, 122,
// 	121, 120, 115, 197, 116, 117, 118, 119, 110, 111,
// 	112, 113, 114, 0, 0, 122, 0, 0, 0, 121,
// 	120, 115, 197, 116, 117, 118, 119, 110, 111, 112,
// 	113, 114, 0, 0, 122, 121, 120, 115, 197, 116,
// 	117, 118, 119, 110, 111, 112, 113, 114, 0, 0,
// 	122, 0, 121, 120, 115, 197, 116, 117, 118, 119,
// 	110, 111, 112, 113, 114, 0, 0, 122, 121, 120,
// 	115, 197, 116, 117, 118, 119, 110, 111, 112, 113,
// 	114, 47, 8, 122, 8, 0, 0, 0, 8, 0,
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 0, 0, 0, 0, 0, 0, 8, 0,
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// 	0, 8, 104, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 0, 0, 8, 0, 0, 0, 0, 0,
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 0, 8,
// }
// var yyPact = [...]int{

// 	860, -1000, 948, -1000, -1000, 16, 948, -51, -25, -1000,
// 	522, 212, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
// 	-1000, -1000, -1000, -1000, 250, -1000, 948, -1000, -11, 239,
// 	196, 511, 149, -1000, 150, -1000, -68, -1000, -1000, 832,
// 	339, -1000, 1197, 92, 11, -1000, -11, -1000, 327, 327,
// 	592, 166, -14, 1077, 470, 327, -69, -70, -71, -72,
// 	-73, -74, -76, -77, -79, -90, 247, -1000, 113, -1000,
// 	55, -1000, 430, 430, 430, 430, 430, 430, 430, 430,
// 	430, 149, 801, 327, 239, -1000, -1000, 167, 250, 32,
// 	-1000, 257, 1445, 41, 327, 166, -1000, -1000, 250, -1000,
// 	-1000, 948, 69, -42, -25, 1221, -1000, 327, 592, 592,
// 	327, 327, 327, 327, 327, 327, 327, 327, 327, 327,
// 	327, 327, 327, -1000, 719, 661, -1000, -45, -45, -93,
// 	58, 885, -1000, -1000, -1000, -1000, -1000, -1000, 263, 112,
// 	-57, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 98,
// 	-60, -1000, -45, 327, 327, 327, 327, 327, 327, -110,
// 	327, 327, 327, -61, 309, 238, -1000, -1000, -1000, -1000,
// 	176, -1000, 327, 1445, 327, 592, 592, 592, 592, 580,
// 	-1000, -1000, -1000, 327, 327, 650, 327, -1000, 250, -1000,
// 	1445, -1000, -1000, -1000, 327, 327, 240, 327, 250, 250,
// 	1172, -81, -1000, -1000, 1445, 33, -5, 35, -43, -43,
// 	-45, -45, -45, -26, -26, -26, -26, -26, -83, 684,
// 	-45, 1316, 327, 166, 1297, 327, 166, -1000, 202, -1000,
// 	-1000, -1000, -1000, 244, -1000, 233, 1147, 1122, 1097, 1072,
// 	1278, 1047, -1000, 1259, 1240, 1022, 242, -1000, -62, -1000,
// 	-80, -1000, 1445, 1445, 3, 3, 3, 3, 247, 226,
// 	3, 1445, 1445, 1445, -14, 1445, -1000, 1429, 394, -1000,
// 	-26, -1000, -1000, -1000, 327, 592, 592, 327, 910, 327,
// 	117, -111, -21, 309, 238, -1000, -1000, -1000, -1000, -1000,
// 	-1000, -1000, -1000, 327, -1000, 327, 327, -1000, 223, 203,
// 	91, 430, 327, 327, -84, 1445, 50, 3, 1445, 327,
// 	1445, 327, -1000, 997, 972, 947, -1000, 1412, 1396, -1000,
// 	327, -1000, 831, 802, -1000, -1000, -1000, 87, -1000, 87,
// 	-1000, 1445, -1000, -1000, 327, -1000, -1000, -1000, -1000, 327,
// 	1377, 1361, -1000, -1000,
// }
// var yyPgo = [...]int{

// 	0, 0, 332, 1551, 330, 142, 1, 329, 328, 327,
// 	326, 325, 267, 266, 29, 324, 311, 23, 5, 282,
// 	3, 4, 2, 298, 295, 288, 147, 67, 286, 284,
// }
// var yyR1 = [...]int{

// 	0, 23, 23, 23, 13, 13, 12, 12, 12, 12,
// 	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
// 	12, 24, 24, 24, 24, 3, 10, 25, 25, 26,
// 	26, 26, 9, 9, 9, 9, 8, 8, 2, 2,
// 	2, 4, 6, 6, 6, 6, 6, 11, 16, 16,
// 	16, 16, 16, 16, 16, 16, 16, 16, 28, 16,
// 	15, 27, 27, 29, 29, 29, 29, 29, 29, 29,
// 	29, 29, 29, 29, 29, 29, 29, 29, 29, 29,
// 	29, 29, 29, 29, 29, 29, 29, 29, 29, 29,
// 	19, 19, 20, 20, 20, 5, 5, 5, 7, 7,
// 	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
// 	14, 14, 17, 17, 17, 17, 17, 17, 17, 17,
// 	17, 17, 17, 17, 17, 18, 18, 18, 21, 21,
// 	21, 22, 22, 22, 22, 22, 22, 22, 22, 1,
// 	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
// 	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
// 	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
// 	1, 1, 1, 1, 1, 1, 1, 1,
// }
// var yyR2 = [...]int{

// 	0, 1, 0, 1, 1, 2, 2, 3, 3, 4,
// 	4, 2, 1, 3, 3, 3, 3, 1, 1, 1,
// 	1, 0, 1, 2, 3, 3, 2, 1, 2, 1,
// 	2, 2, 10, 7, 10, 7, 4, 3, 1, 3,
// 	3, 1, 1, 1, 1, 1, 0, 1, 2, 2,
// 	2, 2, 2, 2, 2, 2, 2, 1, 0, 5,
// 	1, 2, 0, 2, 1, 1, 2, 1, 2, 2,
// 	2, 2, 2, 3, 4, 2, 1, 1, 1, 2,
// 	1, 2, 1, 2, 1, 2, 2, 1, 1, 1,
// 	1, 2, 1, 2, 2, 1, 4, 6, 1, 3,
// 	1, 3, 3, 5, 5, 7, 7, 3, 3, 5,
// 	6, 5, 1, 2, 2, 1, 2, 3, 3, 2,
// 	3, 3, 1, 2, 2, 4, 4, 3, 2, 2,
// 	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
// 	1, 1, 3, 3, 3, 3, 3, 2, 2, 3,
// 	2, 2, 2, 2, 2, 3, 4, 4, 3, 3,
// 	3, 3, 3, 3, 3, 3, 2, 4, 4, 3,
// 	4, 4, 6, 4, 3, 6, 6, 4,
// }
// var yyChk = [...]int{

// 	-1000, -23, -13, 2, -12, -16, -11, 27, -3, 85,
// 	17, 18, -10, -9, -8, 91, 4, 7, 8, 9,
// 	5, 6, 10, 14, -19, 13, -15, 112, 28, 23,
// 	21, 22, -20, 114, -5, 12, 29, -12, 91, -13,
// 	110, 91, -1, -14, -5, 74, 28, -3, 102, 101,
// 	116, -17, 27, -21, 55, 107, 75, 76, 77, 78,
// 	79, 80, 81, 83, 82, 84, 52, 53, -18, 54,
// 	-24, 28, -27, -27, -27, -27, -27, -27, -27, -27,
// 	-27, -20, -13, 92, -25, -26, -5, 19, 20, -4,
// 	28, -2, -1, -5, 116, -17, 32, 32, 116, 109,
// 	-12, 91, -14, 27, -3, -1, 91, 111, 96, 57,
// 	101, 102, 103, 104, 105, 95, 97, 98, 99, 100,
// 	94, 93, 108, 91, 101, 102, 91, -1, -1, -14,
// 	-17, -1, 69, 70, 71, 72, 73, 52, 115, -22,
// 	11, 54, 4, 7, 8, 9, 5, 6, 10, -22,
// 	11, 54, -1, 116, 116, 116, 116, 116, 116, 116,
// 	116, 116, 116, 27, -21, 55, -18, 52, 91, 28,
// 	111, -29, 31, -1, 85, 37, 38, 39, 40, 41,
// 	56, 32, 42, 86, 87, 88, 89, 90, -19, 113,
// 	-1, -26, 30, -5, 37, 92, 24, 96, 99, 100,
// 	-1, -5, -12, 91, -1, -14, -1, -14, -1, -1,
// 	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
// 	-1, -1, 116, -17, -1, 116, -17, 117, 111, 117,
// 	27, 28, 52, 115, 52, 115, -1, -1, -1, -1,
// 	-1, -1, 117, -1, -1, -1, 115, -22, 11, -22,
// 	11, 28, -1, -1, -14, -14, -14, -14, 52, 115,
// 	-14, -1, -1, -1, 27, -1, -28, -1, -1, 25,
// 	-1, -5, -5, 117, 111, 111, 58, 111, -1, 111,
// 	-1, -17, 27, -21, 55, 27, 28, 27, 28, 117,
// 	117, 117, 117, 111, 117, 111, 111, 117, 115, 115,
// 	27, -27, 38, 38, -7, -1, -14, -14, -1, 111,
// 	-1, 111, 117, -1, -1, -1, 52, -1, -1, 117,
// 	111, 95, -1, -1, 117, 117, 117, 40, 26, 40,
// 	26, -1, 117, 117, -6, 101, 102, 103, 104, -6,
// 	-1, -1, 26, 26,
// }
// var yyDef = [...]int{

// 	-2, -2, 1, 3, 4, 0, 0, 0, 0, 12,
// 	0, 21, 17, 18, 19, 20, 62, 62, 62, 62,
// 	62, 62, 62, 62, 62, 57, 0, 47, 0, 0,
// 	0, 0, 90, 60, 92, 95, 0, 5, 6, 0,
// 	0, 11, 0, 0, 0, 139, 140, 141, 0, 0,
// 	0, 100, 112, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 0, 0, 0, 0, 0, 115, 122, 130,
// 	0, 22, 48, 49, 50, 51, 52, 53, 54, 55,
// 	56, 91, 0, 0, 26, 27, 29, 0, 0, 0,
// 	41, 0, 38, 0, 0, 0, 94, 93, 0, 7,
// 	8, 20, 0, 112, 141, 0, 13, 0, 0, 0,
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 0, 14, 0, 0, 15, 147, 148, 0,
// 	100, 0, 150, 151, 152, 153, 154, 113, 0, 116,
// 	138, 128, 131, 132, 133, 134, 135, 136, 137, 119,
// 	138, 129, 166, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 0, 114, 0, 0, 124, 123, 16, 23,
// 	0, 61, 64, 65, 67, 0, 0, 0, 0, 0,
// 	76, 77, 78, 80, 82, 84, 87, 88, 89, 58,
// 	25, 28, 30, 31, 0, 0, 37, 0, 0, 0,
// 	0, 0, 9, 10, 102, 0, -2, 0, 142, 143,
// 	144, 145, 146, -2, -2, -2, -2, -2, 164, 165,
// 	169, 0, 0, 107, 0, 0, 108, 101, 0, 149,
// 	127, 155, 117, 0, 120, 0, 0, 0, 0, 0,
// 	0, 0, 174, 0, 0, 0, 0, 118, 138, 121,
// 	138, 24, 63, 66, 68, 69, 70, 71, 72, 0,
// 	75, 79, 81, 83, -2, 86, 62, 0, 0, 36,
// 	-2, 39, 40, 96, 0, 0, 0, 0, 0, 0,
// 	0, 0, 112, 0, 0, 125, 156, 126, 157, 167,
// 	168, 170, 171, 0, 173, 0, 0, 177, 0, 0,
// 	73, 59, 0, 0, 0, 98, 0, 111, 103, 0,
// 	104, 0, 109, 0, 0, 0, 74, 0, 0, 97,
// 	0, 110, 0, 0, 172, 175, 176, 46, 33, 46,
// 	35, 99, 105, 106, 0, 42, 43, 44, 45, 0,
// 	0, 0, 32, 34,
// }
// var yyTok1 = [...]int{

// 	1, 4, 5, 6, 7, 8, 9, 10, 11, 12,
// 	13, 14, 15, 16, 3, 3, 3, 3, 3, 3,
// 	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
// 	3, 3, 3, 3, 3, 3, 3, 105, 3, 3,
// 	116, 117, 103, 101, 111, 102, 115, 104, 3, 3,
// 	3, 3, 3, 3, 3, 3, 3, 3, 110, 3,
// 	3, 92, 3, 3, 3, 3, 3, 3, 3, 3,
// 	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
// 	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
// 	3, 114, 3, 113, 108, 3, 3, 3, 3, 3,
// 	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
// 	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
// 	3, 3, 3, 112, 3, 109,
// }
// var yyTok2 = [...]int{

// 	2, 3, 0, 0, 0, 0, 0, 0, 0, 0,
// 	0, 0, 0, 0, 0, 17, 18, 19, 20, 21,
// 	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
// 	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
// 	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
// 	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
// 	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
// 	72, 73, 74, 75, 76, 77, 78, 79, 80, 81,
// 	82, 83, 84, 85, 86, 87, 88, 89, 90, 91,
// 	93, 94, 95, 96, 97, 98, 99, 100, 106, 107,
// }
// var yyTok3 = [...]int{
// 	0,
// }

// var yyErrorMessages = [...]struct {
// 	state int
// 	token int
// 	msg   string
// }{
// }

// /*	parser for yacc output	*/

// var (
// 	yyDebug        = 0
// 	yyErrorVerbose = false
// )

// type yyLexer interface {
// 	Lex(lval *yySymType) int
// 	Error(s string)
// }

// type yyParser interface {
// 	Parse(yyLexer) int
// 	Lookahead() int
// }

// type yyParserImpl struct {
// 	lval  yySymType
// 	stack [yyInitialStackSize]yySymType
// 	char  int
// }

// func (p *yyParserImpl) Lookahead() int {
// 	return p.char
// }

// func yyNewParser() yyParser {
// 	return &yyParserImpl{}
// }

// const yyFlag = -1000

// func yyTokname(c int) string {
// 	if c >= 1 && c-1 < len(yyToknames) {
// 		if yyToknames[c-1] != "" {
// 			return yyToknames[c-1]
// 		}
// 	}
// 	return __yyfmt__.Sprintf("tok-%v", c)
// }

// func yyStatname(s int) string {
// 	if s >= 0 && s < len(yyStatenames) {
// 		if yyStatenames[s] != "" {
// 			return yyStatenames[s]
// 		}
// 	}
// 	return __yyfmt__.Sprintf("state-%v", s)
// }

// func yyErrorMessage(state, lookAhead int) string {
// 	const TOKSTART = 4

// 	if !yyErrorVerbose {
// 		return "syntax error"
// 	}

// 	for _, e := range yyErrorMessages {
// 		if e.state == state && e.token == lookAhead {
// 			return "syntax error: " + e.msg
// 		}
// 	}

// 	res := "syntax error: unexpected " + yyTokname(lookAhead)

// 	// To match Bison, suggest at most four expected tokens.
// 	expected := make([]int, 0, 4)

// 	// Look for shiftable tokens.
// 	base := yyPact[state]
// 	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
// 		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
// 			if len(expected) == cap(expected) {
// 				return res
// 			}
// 			expected = append(expected, tok)
// 		}
// 	}

// 	if yyDef[state] == -2 {
// 		i := 0
// 		for yyExca[i] != -1 || yyExca[i+1] != state {
// 			i += 2
// 		}

// 		// Look for tokens that we accept or reduce.
// 		for i += 2; yyExca[i] >= 0; i += 2 {
// 			tok := yyExca[i]
// 			if tok < TOKSTART || yyExca[i+1] == 0 {
// 				continue
// 			}
// 			if len(expected) == cap(expected) {
// 				return res
// 			}
// 			expected = append(expected, tok)
// 		}

// 		// If the default action is to accept or reduce, give up.
// 		if yyExca[i+1] != 0 {
// 			return res
// 		}
// 	}

// 	for i, tok := range expected {
// 		if i == 0 {
// 			res += ", expecting "
// 		} else {
// 			res += " or "
// 		}
// 		res += yyTokname(tok)
// 	}
// 	return res
// }

// func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
// 	token = 0
// 	char = lex.Lex(lval)
// 	if char <= 0 {
// 		token = yyTok1[0]
// 		goto out
// 	}
// 	if char < len(yyTok1) {
// 		token = yyTok1[char]
// 		goto out
// 	}
// 	if char >= yyPrivate {
// 		if char < yyPrivate+len(yyTok2) {
// 			token = yyTok2[char-yyPrivate]
// 			goto out
// 		}
// 	}
// 	for i := 0; i < len(yyTok3); i += 2 {
// 		token = yyTok3[i+0]
// 		if token == char {
// 			token = yyTok3[i+1]
// 			goto out
// 		}
// 	}

// out:
// 	if token == 0 {
// 		token = yyTok2[1] /* unknown char */
// 	}
// 	if yyDebug >= 3 {
// 		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
// 	}
// 	return char, token
// }

// func yyParse(yylex yyLexer) int {
// 	return yyNewParser().Parse(yylex)
// }

// func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
// 	var yyn int
// 	var yyVAL yySymType
// 	var yyDollar []yySymType
// 	_ = yyDollar // silence set and not used
// 	yyS := yyrcvr.stack[:]

// 	Nerrs := 0   /* number of errors */
// 	Errflag := 0 /* error recovery flag */
// 	yystate := 0
// 	yyrcvr.char = -1
// 	yytoken := -1 // yyrcvr.char translated into internal numbering
// 	defer func() {
// 		// Make sure we report no lookahead when not parsing.
// 		yystate = -1
// 		yyrcvr.char = -1
// 		yytoken = -1
// 	}()
// 	yyp := -1
// 	goto yystack

// ret0:
// 	return 0

// ret1:
// 	return 1

// yystack:
// 	/* put a state and value onto the stack */
// 	if yyDebug >= 4 {
// 		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
// 	}

// 	yyp++
// 	if yyp >= len(yyS) {
// 		nyys := make([]yySymType, len(yyS)*2)
// 		copy(nyys, yyS)
// 		yyS = nyys
// 	}
// 	yyS[yyp] = yyVAL
// 	yyS[yyp].yys = yystate

// yynewstate:
// 	yyn = yyPact[yystate]
// 	if yyn <= yyFlag {
// 		goto yydefault /* simple state */
// 	}
// 	if yyrcvr.char < 0 {
// 		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
// 	}
// 	yyn += yytoken
// 	if yyn < 0 || yyn >= yyLast {
// 		goto yydefault
// 	}
// 	yyn = yyAct[yyn]
// 	if yyChk[yyn] == yytoken { /* valid shift */
// 		yyrcvr.char = -1
// 		yytoken = -1
// 		yyVAL = yyrcvr.lval
// 		yystate = yyn
// 		if Errflag > 0 {
// 			Errflag--
// 		}
// 		goto yystack
// 	}

// yydefault:
// 	/* default state action */
// 	yyn = yyDef[yystate]
// 	if yyn == -2 {
// 		if yyrcvr.char < 0 {
// 			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
// 		}

// 		/* look through exception table */
// 		xi := 0
// 		for {
// 			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
// 				break
// 			}
// 			xi += 2
// 		}
// 		for xi += 2; ; xi += 2 {
// 			yyn = yyExca[xi+0]
// 			if yyn < 0 || yyn == yytoken {
// 				break
// 			}
// 		}
// 		yyn = yyExca[xi+1]
// 		if yyn < 0 {
// 			goto ret0
// 		}
// 	}
// 	if yyn == 0 {
// 		/* error ... attempt to resume parsing */
// 		switch Errflag {
// 		case 0: /* brand new error */
// 			yylex.Error(yyErrorMessage(yystate, yytoken))
// 			Nerrs++
// 			if yyDebug >= 1 {
// 				__yyfmt__.Printf("%s", yyStatname(yystate))
// 				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
// 			}
// 			fallthrough

// 		case 1, 2: /* incompletely recovered error ... try again */
// 			Errflag = 3

// 			/* find a state where "error" is a legal shift action */
// 			for yyp >= 0 {
// 				yyn = yyPact[yyS[yyp].yys] + yyErrCode
// 				if yyn >= 0 && yyn < yyLast {
// 					yystate = yyAct[yyn] /* simulate a shift of "error" */
// 					if yyChk[yystate] == yyErrCode {
// 						goto yystack
// 					}
// 				}

// 				/* the current p has no shift on "error", pop stack */
// 				if yyDebug >= 2 {
// 					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
// 				}
// 				yyp--
// 			}
// 			/* there is no state on the stack with an error shift ... abort */
// 			goto ret1

// 		case 3: /* no shift yet; clobber input char */
// 			if yyDebug >= 2 {
// 				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
// 			}
// 			if yytoken == yyEofCode {
// 				goto ret1
// 			}
// 			yyrcvr.char = -1
// 			yytoken = -1
// 			goto yynewstate /* try again in the same state */
// 		}
// 	}

// 	/* reduction by production yyn */
// 	if yyDebug >= 2 {
// 		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
// 	}

// 	yynt := yyn
// 	yypt := yyp
// 	_ = yypt // guard against "declared and not used"

// 	yyp -= yyR2[yyn]
// 	// yyp is now the index of $0. Perform the default action. Iff the
// 	// reduced production is ε, $1 is possibly out of range.
// 	if yyp+1 >= len(yyS) {
// 		nyys := make([]yySymType, len(yyS)*2)
// 		copy(nyys, yyS)
// 		yyS = nyys
// 	}
// 	yyVAL = yyS[yyp+1]

// 	/* consult goto table to find next state */
// 	yyn = yyR1[yyn]
// 	yyg := yyPgo[yyn]
// 	yyj := yyg + yyS[yyp].yys + 1

// 	if yyj >= yyLast {
// 		yystate = yyAct[yyg]
// 	} else {
// 		yystate = yyAct[yyj]
// 		if yyChk[yystate] != -yyn {
// 			yystate = yyAct[yyg]
// 		}
// 	}
// 	// dummy call; replaced with literal code
// 	switch yynt {

// 	case 3:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ WARNING("syntax error"); }
// 	case 6:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ codegen = 1; makeiattr(0, 0); }
// 	case 7:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ rightthing(yyDollar[1].o, '}'); yyVAL.o = yyDollar[2].o; }
// 	case 8:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ y.o=yyDollar[3].o; makevar(yyDollar[1].p,PLACENAME,y); yyVAL.o = yyDollar[3].o; }
// 	case 9:
// 		yyDollar = yyS[yypt-4:yypt+1]
// 		{ y.o=yyDollar[4].o; makevar(yyDollar[1].p,PLACENAME,y); yyVAL.o = yyDollar[4].o; }
// 	case 10:
// 		yyDollar = yyS[yypt-4:yypt+1]
// 		{ y.o=yyDollar[3].o; makevar(yyDollar[1].p,PLACENAME,y); yyVAL.o = yyDollar[3].o; }
// 	case 11:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ y.f = yyDollar[1].f; yyVAL.o = y.o; yyVAL.o = makenode(PLACE, 0); }
// 	case 12:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ setdir(yyDollar[1].i); yyVAL.o = makenode(PLACE, 0); }
// 	case 13:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ printexpr(yyDollar[2].f); yyVAL.o = makenode(PLACE, 0); }
// 	case 14:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ printpos(yyDollar[2].o); yyVAL.o = makenode(PLACE, 0); }
// 	case 15:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ printf("%s\n", yyDollar[2].p); free(yyDollar[2].p); yyVAL.o = makenode(PLACE, 0); }
// 	case 16:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ resetvar(); makeiattr(0, 0); yyVAL.o = makenode(PLACE, 0); }
// 	case 22:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ makevattr(yyDollar[1].p); }
// 	case 23:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ makevattr(yyDollar[2].p); }
// 	case 24:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ makevattr(yyDollar[3].p); }
// 	case 25:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.f=y.f=yyDollar[3].f; makevar(yyDollar[1].p,VARNAME,y); checkscale(yyDollar[1].p); }
// 	case 26:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ copy(); }
// 	case 29:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ copyfile(yyDollar[1].p); }
// 	case 30:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ copydef(yyDollar[2].st); }
// 	case 31:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ copyuntil(yyDollar[2].p); }
// 	case 32:
// 		yyDollar = yyS[yypt-10:yypt+1]
// 		{ forloop(yyDollar[2].p, yyDollar[4].f, yyDollar[6].f, yyDollar[8].i, yyDollar[9].f, yyDollar[10].p); }
// 	case 33:
// 		yyDollar = yyS[yypt-7:yypt+1]
// 		{ forloop(yyDollar[2].p, yyDollar[4].f, yyDollar[6].f, '+', 1.0, yyDollar[7].p); }
// 	case 34:
// 		yyDollar = yyS[yypt-10:yypt+1]
// 		{ forloop(yyDollar[2].p, yyDollar[4].f, yyDollar[6].f, yyDollar[8].i, yyDollar[9].f, yyDollar[10].p); }
// 	case 35:
// 		yyDollar = yyS[yypt-7:yypt+1]
// 		{ forloop(yyDollar[2].p, yyDollar[4].f, yyDollar[6].f, '+', 1.0, yyDollar[7].p); }
// 	case 36:
// 		yyDollar = yyS[yypt-4:yypt+1]
// 		{ ifstat(yyDollar[2].f, yyDollar[3].p, yyDollar[4].p); }
// 	case 37:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ ifstat(yyDollar[2].f, yyDollar[3].p, (char *) 0); }
// 	case 39:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.f = strcmp(yyDollar[1].p,yyDollar[3].p) == 0; free(yyDollar[1].p); free(yyDollar[3].p); }
// 	case 40:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.f = strcmp(yyDollar[1].p,yyDollar[3].p) != 0; free(yyDollar[1].p); free(yyDollar[3].p); }
// 	case 41:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ y.f = 0; makevar(yyDollar[1].p, VARNAME, y); }
// 	case 42:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ yyVAL.i = '+'; }
// 	case 43:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ yyVAL.i = '-'; }
// 	case 44:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ yyVAL.i = '*'; }
// 	case 45:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ yyVAL.i = '/'; }
// 	case 46:
// 		yyDollar = yyS[yypt-0:yypt+1]
// 		{ yyVAL.i = ' '; }
// 	case 47:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ yyVAL.o = leftthing('{'); }
// 	case 48:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.o = boxgen(); }
// 	case 49:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.o = circgen(yyDollar[1].i); }
// 	case 50:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.o = circgen(yyDollar[1].i); }
// 	case 51:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.o = arcgen(yyDollar[1].i); }
// 	case 52:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.o = linegen(yyDollar[1].i); }
// 	case 53:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.o = linegen(yyDollar[1].i); }
// 	case 54:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.o = linegen(yyDollar[1].i); }
// 	case 55:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.o = movegen(); }
// 	case 56:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.o = textgen(); }
// 	case 57:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ yyVAL.o = troffgen(yyDollar[1].p); }
// 	case 58:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.o=rightthing(yyDollar[1].o,']'); }
// 	case 59:
// 		yyDollar = yyS[yypt-5:yypt+1]
// 		{ yyVAL.o = blockgen(yyDollar[1].o, yyDollar[4].o); }
// 	case 60:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ yyVAL.o = leftthing('['); }
// 	case 63:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ makefattr(yyDollar[1].i, !DEFAULT, yyDollar[2].f); }
// 	case 64:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ makefattr(yyDollar[1].i, DEFAULT, 0.0); }
// 	case 65:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ makefattr(curdir(), !DEFAULT, yyDollar[1].f); }
// 	case 66:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ makefattr(yyDollar[1].i, !DEFAULT, yyDollar[2].f); }
// 	case 67:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ makefattr(yyDollar[1].i, DEFAULT, 0.0); }
// 	case 68:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ makeoattr(yyDollar[1].i, yyDollar[2].o); }
// 	case 69:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ makeoattr(yyDollar[1].i, yyDollar[2].o); }
// 	case 70:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ makeoattr(yyDollar[1].i, yyDollar[2].o); }
// 	case 71:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ makeoattr(yyDollar[1].i, yyDollar[2].o); }
// 	case 72:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ makeiattr(WITH, yyDollar[2].i); }
// 	case 73:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ makeoattr(PLACE, getblock(getlast(1,BLOCK), yyDollar[3].p)); }
// 	case 74:
// 		yyDollar = yyS[yypt-4:yypt+1]
// 		{ makeoattr(PLACE, getpos(getblock(getlast(1,BLOCK), yyDollar[3].p), yyDollar[4].i)); }
// 	case 75:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ makeoattr(PLACE, yyDollar[2].o); }
// 	case 76:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ makeiattr(SAME, yyDollar[1].i); }
// 	case 77:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ maketattr(yyDollar[1].i, (char *) 0); }
// 	case 78:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ makeiattr(HEAD, yyDollar[1].i); }
// 	case 79:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ makefattr(DOT, !DEFAULT, yyDollar[2].f); }
// 	case 80:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ makefattr(DOT, DEFAULT, 0.0); }
// 	case 81:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ makefattr(DASH, !DEFAULT, yyDollar[2].f); }
// 	case 82:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ makefattr(DASH, DEFAULT, 0.0); }
// 	case 83:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ makefattr(CHOP, !DEFAULT, yyDollar[2].f); }
// 	case 84:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ makefattr(CHOP, DEFAULT, 0.0); }
// 	case 85:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ makeattr(CHOP, PLACENAME, getvar(yyDollar[2].p)); }
// 	case 86:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ makefattr(FILL, !DEFAULT, yyDollar[2].f); }
// 	case 87:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ makefattr(FILL, DEFAULT, 0.0); }
// 	case 88:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ makeiattr(NOEDGE, 0); }
// 	case 92:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ maketattr(CENTER, yyDollar[1].p); }
// 	case 93:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ maketattr(yyDollar[2].i, yyDollar[1].p); }
// 	case 94:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ addtattr(yyDollar[2].i); }
// 	case 96:
// 		yyDollar = yyS[yypt-4:yypt+1]
// 		{ yyVAL.p = sprintgen(yyDollar[3].p); }
// 	case 97:
// 		yyDollar = yyS[yypt-6:yypt+1]
// 		{ yyVAL.p = sprintgen(yyDollar[3].p); }
// 	case 98:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ exprsave(yyDollar[1].f); yyVAL.i = 0; }
// 	case 99:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ exprsave(yyDollar[3].f); }
// 	case 101:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.o = yyDollar[2].o; }
// 	case 102:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.o = makepos(yyDollar[1].f, yyDollar[3].f); }
// 	case 103:
// 		yyDollar = yyS[yypt-5:yypt+1]
// 		{ yyVAL.o = fixpos(yyDollar[1].o, yyDollar[3].f, yyDollar[5].f); }
// 	case 104:
// 		yyDollar = yyS[yypt-5:yypt+1]
// 		{ yyVAL.o = fixpos(yyDollar[1].o, -yyDollar[3].f, -yyDollar[5].f); }
// 	case 105:
// 		yyDollar = yyS[yypt-7:yypt+1]
// 		{ yyVAL.o = fixpos(yyDollar[1].o, yyDollar[4].f, yyDollar[6].f); }
// 	case 106:
// 		yyDollar = yyS[yypt-7:yypt+1]
// 		{ yyVAL.o = fixpos(yyDollar[1].o, -yyDollar[4].f, -yyDollar[6].f); }
// 	case 107:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.o = addpos(yyDollar[1].o, yyDollar[3].o); }
// 	case 108:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.o = subpos(yyDollar[1].o, yyDollar[3].o); }
// 	case 109:
// 		yyDollar = yyS[yypt-5:yypt+1]
// 		{ yyVAL.o = makepos(getcomp(yyDollar[2].o,DOTX), getcomp(yyDollar[4].o,DOTY)); }
// 	case 110:
// 		yyDollar = yyS[yypt-6:yypt+1]
// 		{ yyVAL.o = makebetween(yyDollar[1].f, yyDollar[3].o, yyDollar[5].o); }
// 	case 111:
// 		yyDollar = yyS[yypt-5:yypt+1]
// 		{ yyVAL.o = makebetween(yyDollar[1].f, yyDollar[3].o, yyDollar[5].o); }
// 	case 112:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ y = getvar(yyDollar[1].p); yyVAL.o = y.o; }
// 	case 113:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ y = getvar(yyDollar[1].p); yyVAL.o = getpos(y.o, yyDollar[2].i); }
// 	case 114:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ y = getvar(yyDollar[2].p); yyVAL.o = getpos(y.o, yyDollar[1].i); }
// 	case 115:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ yyVAL.o = gethere(); }
// 	case 116:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.o = getlast(yyDollar[1].i, yyDollar[2].i); }
// 	case 117:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.o = getpos(getlast(yyDollar[1].i, yyDollar[2].i), yyDollar[3].i); }
// 	case 118:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.o = getpos(getlast(yyDollar[2].i, yyDollar[3].i), yyDollar[1].i); }
// 	case 119:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.o = getfirst(yyDollar[1].i, yyDollar[2].i); }
// 	case 120:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.o = getpos(getfirst(yyDollar[1].i, yyDollar[2].i), yyDollar[3].i); }
// 	case 121:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.o = getpos(getfirst(yyDollar[2].i, yyDollar[3].i), yyDollar[1].i); }
// 	case 123:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.o = getpos(yyDollar[1].o, yyDollar[2].i); }
// 	case 124:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.o = getpos(yyDollar[2].o, yyDollar[1].i); }
// 	case 125:
// 		yyDollar = yyS[yypt-4:yypt+1]
// 		{ yyVAL.o = getblock(getlast(yyDollar[1].i,yyDollar[2].i), yyDollar[4].p); }
// 	case 126:
// 		yyDollar = yyS[yypt-4:yypt+1]
// 		{ yyVAL.o = getblock(getfirst(yyDollar[1].i,yyDollar[2].i), yyDollar[4].p); }
// 	case 127:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ y = getvar(yyDollar[1].p); yyVAL.o = getblock(y.o, yyDollar[3].p); }
// 	case 128:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.i = yyDollar[1].i + 1; }
// 	case 129:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.i = yyDollar[1].i; }
// 	case 130:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ yyVAL.i = 1; }
// 	case 140:
// 		yyDollar = yyS[yypt-1:yypt+1]
// 		{ yyVAL.f = getfval(yyDollar[1].p); }
// 	case 142:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.f = yyDollar[1].f + yyDollar[3].f; }
// 	case 143:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.f = yyDollar[1].f - yyDollar[3].f; }
// 	case 144:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.f = yyDollar[1].f * yyDollar[3].f; }
// 	case 145:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ if (yyDollar[3].f == 0.0) {
// 						WARNING("division by 0"); yyDollar[3].f = 1; }
// 					  yyVAL.f = yyDollar[1].f / yyDollar[3].f; }
// 	case 146:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ if ((long)yyDollar[3].f == 0) {
// 						WARNING("mod division by 0"); yyDollar[3].f = 1; }
// 					  yyVAL.f = (long)yyDollar[1].f % (long)yyDollar[3].f; }
// 	case 147:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.f = -yyDollar[2].f; }
// 	case 148:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.f = yyDollar[2].f; }
// 	case 149:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.f = yyDollar[2].f; }
// 	case 150:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.f = getcomp(yyDollar[1].o, yyDollar[2].i); }
// 	case 151:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.f = getcomp(yyDollar[1].o, yyDollar[2].i); }
// 	case 152:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.f = getcomp(yyDollar[1].o, yyDollar[2].i); }
// 	case 153:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.f = getcomp(yyDollar[1].o, yyDollar[2].i); }
// 	case 154:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.f = getcomp(yyDollar[1].o, yyDollar[2].i); }
// 	case 155:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ y = getvar(yyDollar[1].p); yyVAL.f = getblkvar(y.o, yyDollar[3].p); }
// 	case 156:
// 		yyDollar = yyS[yypt-4:yypt+1]
// 		{ yyVAL.f = getblkvar(getlast(yyDollar[1].i,yyDollar[2].i), yyDollar[4].p); }
// 	case 157:
// 		yyDollar = yyS[yypt-4:yypt+1]
// 		{ yyVAL.f = getblkvar(getfirst(yyDollar[1].i,yyDollar[2].i), yyDollar[4].p); }
// 	case 158:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.f = yyDollar[1].f > yyDollar[3].f; }
// 	case 159:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.f = yyDollar[1].f < yyDollar[3].f; }
// 	case 160:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.f = yyDollar[1].f <= yyDollar[3].f; }
// 	case 161:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.f = yyDollar[1].f >= yyDollar[3].f; }
// 	case 162:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.f = yyDollar[1].f == yyDollar[3].f; }
// 	case 163:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.f = yyDollar[1].f != yyDollar[3].f; }
// 	case 164:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.f = yyDollar[1].f && yyDollar[3].f; }
// 	case 165:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.f = yyDollar[1].f || yyDollar[3].f; }
// 	case 166:
// 		yyDollar = yyS[yypt-2:yypt+1]
// 		{ yyVAL.f = !(yyDollar[2].f); }
// 	case 167:
// 		yyDollar = yyS[yypt-4:yypt+1]
// 		{ yyVAL.f = Log10(yyDollar[3].f); }
// 	case 168:
// 		yyDollar = yyS[yypt-4:yypt+1]
// 		{ yyVAL.f = Exp(yyDollar[3].f * log(10.0)); }
// 	case 169:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.f = pow(yyDollar[1].f, yyDollar[3].f); }
// 	case 170:
// 		yyDollar = yyS[yypt-4:yypt+1]
// 		{ yyVAL.f = sin(yyDollar[3].f); }
// 	case 171:
// 		yyDollar = yyS[yypt-4:yypt+1]
// 		{ yyVAL.f = cos(yyDollar[3].f); }
// 	case 172:
// 		yyDollar = yyS[yypt-6:yypt+1]
// 		{ yyVAL.f = atan2(yyDollar[3].f, yyDollar[5].f); }
// 	case 173:
// 		yyDollar = yyS[yypt-4:yypt+1]
// 		{ yyVAL.f = Sqrt(yyDollar[3].f); }
// 	case 174:
// 		yyDollar = yyS[yypt-3:yypt+1]
// 		{ yyVAL.f = (float)rand() / RAND_MAX; }
// 	case 175:
// 		yyDollar = yyS[yypt-6:yypt+1]
// 		{ yyVAL.f = yyDollar[3].f >= yyDollar[5].f ? yyDollar[3].f : yyDollar[5].f; }
// 	case 176:
// 		yyDollar = yyS[yypt-6:yypt+1]
// 		{ yyVAL.f = yyDollar[3].f <= yyDollar[5].f ? yyDollar[3].f : yyDollar[5].f; }
// 	case 177:
// 		yyDollar = yyS[yypt-4:yypt+1]
// 		{ yyVAL.f = (long) yyDollar[3].f; }
// 	}
// 	goto yystack /* stack new state and value */
// }
