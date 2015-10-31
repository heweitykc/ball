module game {
    
    export class CSCommand {
        public static START: string = "1";
        public static END: string = "6";
        public static HELLO: string = "100";
        public static UP: string = "2";
        public static DOWN: string = "3";
        public static RIGHT: string = "4";
        public static LEFT: string = "5";
    
    }
    //服务端通讯
    export class NetCommand extends puremvc.SimpleCommand implements puremvc.ICommand {
        private client: NetClient;
        public constructor() {
            super();
        }

        private msgCallback(msg: string): void {
            console.log(msg);      
            if(msg == CSCommand.START){
                CommonData.isRunning = true;
            } else if(msg == CSCommand.HELLO) {
                CommonData.isRunning = true;
            }
        }

        public register(): void {                        
            this.client = new NetClient();
            this.client.msgCallback = this.msgCallback;
            this.client.Connect();
            CommonData.isRunning = false;
        }
    }
}