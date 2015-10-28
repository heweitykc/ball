

class Main extends egret.Sprite{

    /**
     * 加载进度界面
     */
    private loadingView:LoadingUI;

    public constructor() {
        super();
        this.addEventListener(egret.Event.ADDED_TO_STAGE, this.init, this)
    }

    private appContainer:game.AppContainer;
    private client:NetClient;

    private init(event:egret.Event):void {
        //设置自定义的屏幕适配方式
        egret.sys.screenAdapter = new AutoScreenAdapter();

        //注入自定义的解析器
        egret.gui.mapClass("egret.gui.IAssetAdapter", AssetAdapter);
        egret.gui.Theme.load("resource/theme.thm");

        //初始化UIStage
        this.appContainer = new game.AppContainer();
        this.addChild(this.appContainer);

        //初始化Resource资源加载库
        RES.addEventListener(RES.ResourceEvent.CONFIG_COMPLETE, this.onConfigComplete, this);
        RES.loadConfig("resource/resource.json", "resource/");
        
        this.client = new NetClient();
        this.client.Connect();
    }

    /**
     * 配置文件加载完成,开始预加载preload资源组。
     */
    private onConfigComplete(event:RES.ResourceEvent):void{
        RES.removeEventListener(RES.ResourceEvent.CONFIG_COMPLETE,this.onConfigComplete,this);
        RES.addEventListener(RES.ResourceEvent.GROUP_COMPLETE,this.onResourceLoadComplete,this);
        RES.addEventListener(RES.ResourceEvent.GROUP_PROGRESS,this.onResourceProgress,this);
        RES.loadGroup("loading",1);
        RES.loadGroup("preload");
    }
    /**
     * preload资源组加载完成
     */
    private onResourceLoadComplete(event:RES.ResourceEvent):void {
        if(event.groupName=="loading"){
            //设置加载进度界面
            this.loadingView  = new LoadingUI();
            this.appContainer.addElement(this.loadingView);
        }
        else if(event.groupName=="preload"){
            this.appContainer.removeElement(this.loadingView);
            RES.removeEventListener(RES.ResourceEvent.GROUP_COMPLETE,this.onResourceLoadComplete,this);
            RES.removeEventListener(RES.ResourceEvent.GROUP_PROGRESS,this.onResourceProgress,this);
            this.createGameScene();
        }
    }
    /**
     * preload资源组加载进度
     */
    private onResourceProgress(event:RES.ResourceEvent):void {
        if(event.groupName=="preload"){
            this.loadingView.setProgress(event.itemsLoaded,event.itemsTotal);
        }
    }

    /**
     * 创建游戏场景
     */
    private createGameScene():void {
        //设置模态层透明度,写在这里是因为初始化的时候UIStage还没初始化完毕，直接设置会报错
        egret.gui.PopUpManager.modalAlpha = 0;

        game.ApplicationFacade.getInstance().startUp(this.appContainer);
        game.ApplicationFacade.getInstance().sendNotification(game.SceneCommand.CHANGE,1);
    }
}