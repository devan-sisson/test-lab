"use strict";
var magicNum;
(function (magicNum) {
    magicNum["a"] = "a";
    magicNum["b"] = "b";
})(magicNum || (magicNum = {}));
var magic;
(function (magic) {
    magic[magic["a"] = 0] = "a";
    magic[magic["b"] = 1] = "b";
    magic[magic["c"] = 2] = "c";
    magic[magic["d"] = 3] = "d";
})(magic || (magic = {}));
