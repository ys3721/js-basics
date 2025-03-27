import * as Phaser from "phaser";

export default class GameScene extends Phaser.Scene {
    private socket!: WebSocket;

    constructor() {
        super({ key: "GameScene" });
    }

    preload() {
        // 加载资源
    }

    create() {
        // 初始化 WebSocket 连接到后端
        this.socket = new WebSocket("ws://localhost:8080/ws");

        this.socket.onopen = () => {
            console.log("WebSocket connected");
            // 发送初始消息
            this.socket.send("Hello from client!");
        };

        this.socket.onmessage = (event) => {
            console.log("Received from server:", event.data);
            // 根据后端返回的数据更新游戏状态
        };

        this.socket.onclose = () => {
            console.log("WebSocket closed");
        };

        // 初始化游戏场景，例如添加角色、地图等
        this.add.text(100, 100, "Real-time Game", { font: "24px Arial" });
    }

    update(time: number, delta: number) {
        // 游戏更新逻辑，比如玩家输入处理和状态同步
    }
}
