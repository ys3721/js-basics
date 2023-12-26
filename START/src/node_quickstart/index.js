const { MongoClient } = require("mongodb");

const uri = "mongodb://127.0.0.1:27017/?maxIdleTimeMS=30000"

const client = new MongoClient(uri);

async function run() {
    try {
        const database = client.db('ascension_debug');
        const codes = database.collection('gift_code_batch');

        const query = {operator:"谢经松 新vip礼包1014"};
        const role = await codes.findOne(query);

        console.log(role)
    } finally {
        await client.close()
    }
}


async function findcodes(operatorInfo) {
    try {
        const database = client.db('ascension_debug');
        const codes = database.collection('gift_code_batch');

        const query = {operator:operatorInfo};
        const role = await codes.findOne(query);

        console.log(role)
    } finally {
        await client.close()
    }
}
