/**
 *
 * @author 
 *
 */
class NetClient {
    private socket: egret.WebSocket;
    public msgCallback: Function;
	public constructor() {
        this.socket = new egret.WebSocket();
        this.socket.type = egret.WebSocket.TYPE_STRING;
        this.socket.addEventListener(egret.ProgressEvent.SOCKET_DATA,this.onReceiveMessage,this);
        this.socket.addEventListener(egret.Event.CONNECT,this.onSocketOpen,this);        
        this.socket.addEventListener(egret.Event.CLOSE,this.onSocketClose,this);
	}
	
    public Connect(): void {
        //this.socket.connectByUrl("ws://182.254.149.35:1234/socket");
        this.socket.connectByUrl("ws://127.0.0.1:1234/socket");
        console.log("socket start");
    }
    
    private onReceiveMessage(evt: egret.ProgressEvent): void {
        console.log("recv data : " + evt.bytesTotal);
        //var bts: egret.ByteArray = new egret.ByteArray();
        //this.socket.readBytes(bts);
    
        var msg = this.socket.readUTF();
        this.msgCallback(msg);
    }
    
    private onSocketOpen(evt: egret.ProgressEvent): void {
        console.log("socket connected");
    }
    
    private onSocketClose(evt: egret.ProgressEvent): void {
        console.log("socket close");
    }
}
