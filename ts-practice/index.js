"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
const mongodb_1 = require("mongodb");
const url = "mongodb://localhost:27017";
const client = new mongodb_1.MongoClient(url);
function sayHello(name) {
    console.log(`Hello, ${name}!`);
}
sayHello('World');
function fetchScoreFromDB() {
    return __awaiter(this, void 0, void 0, function* () {
        try {
            yield client.connect();
            const db = client.db("kof");
            const collection = db.collection("player");
            const count = yield collection.countDocuments();
            return count;
        }
        catch (error) {
            console.error(error);
            return 0; // Return a default value in case of an error
        }
        finally {
            yield client.close();
        }
    });
}
function fetchScores() {
    return new Promise(resolve => {
        setTimeout(() => {
            let score = Math.floor(Math.random() * 100);
            resolve(score);
        }, 1000);
    });
}
function delay(seconds) {
    return new Promise(r => {
        setTimeout(r, seconds * 1000);
    });
}
function updateScoreLoop() {
    return __awaiter(this, void 0, void 0, function* () {
        while (true) {
            try {
                let score = yield fetchScoreFromDB();
                console.log(`分数已经更新:${score}`);
            }
            catch (e) {
                console.error("获取分数失败", e);
            }
            yield delay(3); // 延迟3秒再请求下一次
        }
    });
}
updateScoreLoop();
