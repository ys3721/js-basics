var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (Object.prototype.hasOwnProperty.call(b, p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        if (typeof b !== "function" && b !== null)
            throw new TypeError("Class extends value " + String(b) + " is not a constructor or null");
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
import * as Phaser from "phaser";
var GameScene = /** @class */ (function (_super) {
    __extends(GameScene, _super);
    function GameScene() {
        return _super.call(this, { key: "GameScene" }) || this;
    }
    GameScene.prototype.preload = function () {
        // 加载资源
    };
    GameScene.prototype.create = function () {
        var _this = this;
        // 初始化 WebSocket 连接到后端
        this.socket = new WebSocket("ws://localhost:8080/ws");
        this.socket.onopen = function () {
            console.log("WebSocket connected");
            // 发送初始消息
            _this.socket.send("Hello from client!");
        };
        this.socket.onmessage = function (event) {
            console.log("Received from server:", event.data);
            // 根据后端返回的数据更新游戏状态
        };
        this.socket.onclose = function () {
            console.log("WebSocket closed");
        };
        // 初始化游戏场景，例如添加角色、地图等
        this.add.text(100, 100, "Real-time Game", { font: "24px Arial" });
    };
    GameScene.prototype.update = function (time, delta) {
        // 游戏更新逻辑，比如玩家输入处理和状态同步
    };
    return GameScene;
}(Phaser.Scene));
export default GameScene;
