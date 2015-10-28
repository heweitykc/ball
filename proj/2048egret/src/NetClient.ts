/**
 *
 * @author 
 *
 */
class NetClient {
    private socket: egret.WebSocket;
    
	public constructor() {
        this.socket = new egret.WebSocket();
        this.socket.type = egret.WebSocket.TYPE_BINARY;
        this.socket.addEventListener(egret.ProgressEvent.SOCKET_DATA,this.onReceiveMessage,this);
        this.socket.addEventListener(egret.Event.CONNECT,this.onSocketOpen,this);        
        this.socket.addEventListener(egret.Event.CLOSE,this.onSocketClose,this);
	}
	
    public Connect(): void {
        this.socket.connectByUrl("ws://localhost:1234/socket");
        console.log("socket start");
    }
    
    private onReceiveMessage(evt: egret.ProgressEvent): void {
        //var byte: egret.ByteArray = new egret.ByteArray();
        //this.socket.readBytes(byte);
        //var msg: number = byte.readInt();
        console.log("recv data : " + evt.bytesLoaded + ","+evt.bytesTotal);
    }
    private onSocketOpen(evt: egret.ProgressEvent): void {
        console.log("socket connected");
        var byte: egret.ByteArray = new egret.ByteArray();
        byte.writeUTF("Hello Egret WebSocket");
        byte.position = 0;
        this.socket.writeBytes(byte,0,byte.bytesAvailable);
        this.socket.flush();
    }
    private onSocketClose(evt: egret.ProgressEvent): void {
        console.log("socket close");
    }
}
