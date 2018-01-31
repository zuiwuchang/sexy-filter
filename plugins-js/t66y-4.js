(function(){
    //const
    var BaseUrl = 'http://t66y.com/thread0806.php?fid=4&search=&page=';
    var BasePages = 'http://t66y.com/';

    var TrBegin = '<tr class="tr3 t_one tac">';
    var TrEng = '</tr>';
    var H3Begin = '<h3><a href="'
    var H3End = '</h3>'
    var Tag = "</b></a>";
    
    var analyzeTr = function(str){
        var begin = str.indexOf(H3Begin)
        if(begin == -1){
            return;
        }
        begin += H3Begin.length;

        var end = str.indexOf(H3End,begin);
        if(end == -1){
            return;
        }
        str = str.substring(begin,end)
        if(str.endsWith(Tag)){
            return;
        }
        //url
        begin = str.indexOf('"');
        if(begin < 1){
            return;
        }
        var url = str.substring(0,begin)

        //title
        str = str.substring(begin)
        begin = str.indexOf(">");
        if(begin == -1){
            return;
        }
        begin += 1;

        end = str.indexOf("</a>",begin);
        if(end == -1){
            return;
        }
        var title = str.substring(begin,end);

        return {
            Url:BasePages + url,
            Title:title,
        }
    }
    //數據 分析 實現
    var analyze = function(str){
        var rs = [];
        var pos = 0;
        while(true){
            var begin = str.indexOf(TrBegin,pos);
            if(begin == -1){
                break;
            }
            begin += TrBegin.length;

            var end = str.indexOf(TrEng,begin);
            if(end == -1){
                break;
            }

            var node = analyzeTr(str.substring(begin,end));
            if(node){
                rs.push(node);
            }
            pos = end + TrEng.length;
        }
        
        //返回 分析 結果
        return rs;
    };
    //返回 插件 object
    return {
        //返回 插件 唯一 標識
        Id:function(){
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
            
            //返回 分析 結果
            return analyze(str);
        },
        //返回 請求url地址 如果 返回 false 則停止 數據爬取
        //i 是 第多次是 請求 從0 開始計數
        GetUrl:function(i){
            if(i > 99){
                return false;
            }
            return BaseUrl + (i+1);
        },
    }
})();