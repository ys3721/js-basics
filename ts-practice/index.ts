import { MongoClient } from 'mongodb';

const url = "mongodb://localhost:27017";
const client = new MongoClient(url);

function sayHello(name: string) {
    console.log(`Hello, ${name}!`);
}
sayHello('World');

async function fetchScoreFromDB(): Promise<number> {
    try {
        await client.connect();
        const db = client.db("kof")
        const collection = db.collection("player")
        const count = await collection.countDocuments();
        return count;
    }catch (error) {
        console.error(error);
        return 0; // Return a default value in case of an error
    } finally {
        await client.close()
    }
}

function fetchScores(): Promise<number> {
    return new Promise(resolve => {
        setTimeout(() => {
            let score = Math.floor(Math.random()* 100)
            resolve(score)
        }, 1000);
    });
}

function delay(seconds:number) {
    return new Promise<void>(r => {
        setTimeout(r, seconds*1000)
    });
} 

async function updateScoreLoop() {
    while (true) {
        try {
            let score = await fetchScoreFromDB()
            console.log(`分数已经更新:${score}`)
        } catch (e) {
            console.error("获取分数失败", e)
        }
        await delay(3); // 延迟3秒再请求下一次
    }
}

updateScoreLoop();