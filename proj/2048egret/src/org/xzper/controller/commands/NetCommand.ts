module game {
    
    class CSCommand {        
        public static END: string = "6";
        public static HELLO0: string = "100";
        public static HELLO1: string = "101";
        
        public static UP: string = "2";
        public static DOWN: string = "3";
        public static RIGHT: string = "4";
        public static LEFT: string = "5";
        
        
        public static START: string = "8";
        public static TILE: string = "9";
    
    }
    //服务端通讯
    export class NetCommand extends puremvc.SimpleCommand implements puremvc.ICommand {                
        public constructor() {
            super();
        }        
        
        private msgCallback(msg: string): void {
            var gridProxy: GridProxy = <GridProxy><any> (this.facade.retrieveProxy(GridProxy.NAME));
            console.log("server cmd:"+msg);      
            if(msg == CSCommand.HELLO0) {
                CommonData.isRunning = false;
                CommonData.id = 0;
            } else if(msg == CSCommand.HELLO1) {
                CommonData.isRunning = false;
                CommonData.id = 1;
            } else if(msg == CSCommand.UP) {
                gridProxy.move(0);
                CommonData.step++;
            } else if(msg == CSCommand.DOWN) {
                gridProxy.move(2);
                CommonData.step++;
            } else if(msg == CSCommand.RIGHT) {
                gridProxy.move(1);
                CommonData.step++;
            } else if(msg == CSCommand.LEFT) {
                gridProxy.move(3);
                CommonData.step++;
            } else if(msg[0] == CSCommand.START) {                
                //开始游戏        
                var msgbody:string = msg.slice(1,msg.length);
                console.log("开始游戏:" + msgbody);
                var arr = msgbody.split("|");
                for(var i = 0;i < arr.length;i++) {
                    var arr0 = arr[i].split(",");
                    var pos = { x: Number(arr0[0]),y: Number(arr0[1]) };
                    gridProxy.addNetTile(pos,2);
                }
                CommonData.isRunning = true;
            } else if(msg[0] == CSCommand.TILE) {
                //生成tile,解析消息
                var msgbody:string = msg.slice(1,msg.length);
                var newtile: string = msgbody;
                console.log("新tile:" + msgbody);
            }
        }

        public register(): void {                        
            this.Connect();
            CommonData.isRunning = false;            
            this.facade.registerCommand(GameCommand.MOVE_TILE,NetCommand);            
        }               
        
        public execute(notification: puremvc.INotification): void {
            var data: any = notification.getBody();            
            switch(notification.getName()) {
                case GameCommand.MOVE_TILE: { 
                    this.onMove(data);                    
                    break;
                }
            }
        }
        
        private onMove(direction: number): void
        {
            var ret = CommonData.step % 2 ;
            console.log("ret="+ret);
            if(ret != CommonData.id) return;
            if(direction == 0)
                this.SendMsg(CSCommand.UP);
            else if(direction == 1)
                this.SendMsg(CSCommand.RIGHT);
            else if(direction == 2)
                this.SendMsg(CSCommand.DOWN);
            else if(direction == 3)
                this.SendMsg(CSCommand.LEFT);
        }
        
        public Connect(): void {
            CommonData.socket = new egret.WebSocket();
            CommonData.socket.type = egret.WebSocket.TYPE_STRING;
            CommonData.socket.addEventListener(egret.ProgressEvent.SOCKET_DATA,this.onReceiveMessage,this);
            CommonData.socket.addEventListener(egret.Event.CONNECT,this.onSocketOpen,this);
            CommonData.socket.addEventListener(egret.Event.CLOSE,this.onSocketClose,this);
            CommonData.socket.connectByUrl("ws://182.254.149.35:1234/socket");
            //CommonData.socket.connectByUrl("ws://127.0.0.1:1234/socket");
            console.log("socket start");
        }

        private onReceiveMessage(evt: egret.ProgressEvent): void {
            var msg = CommonData.socket.readUTF();
            this.msgCallback(msg);
        }

        public SendMsg(msg: string): void {
            console.log("send:" + msg);
            CommonData.socket.writeUTF(msg);
            CommonData.socket.flush();
        }

        private onSocketOpen(evt: egret.ProgressEvent): void {
            console.log("socket connected");
        }

        private onSocketClose(evt: egret.ProgressEvent): void {
            console.log("socket close");
        }
    }
}