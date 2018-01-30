(function(){
    //返回 插件 object
    return {
        //返回 插件 唯一 標識
        ID:function(){
            return "t66y-west"
        },
        //返回 插件 名稱
        Name:function(){
            return "草榴社區-歐美原創區"
        },
        //數據解析 函數
        Analyze:function(str){
            //驗證 插件 合法性
            if(str == "Analyze"){
                return str;
            }

            //實現 分析結果
            return [
                {
                    Url:"123",
                    Title:"wde",
                },
                {
                    Url:"456",
                    Title:"草泥馬",
                },
            ];
        },
    }
})();