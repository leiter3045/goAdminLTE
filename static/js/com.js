(function($){
    $.com = $.com || {version: "v1.0.0"},
        $.extend($.com, {
            util: {
                getStrLength: function(str) {
                    str = $.trim(str);
                    var length = str.replace(/[^\x00-\xff]/g, "**").length;
                    return parseInt(length / 2) == length / 2 ? length / 2: parseInt(length / 2) + .5;
                },
                isEmpty: function(str) {
                    return void 0 === str || null === str || "" === str
                },
                isInt: function(val){
                    return /^[\d]+$/.test(val);
                },
                isMobile: function(num){
                    return /^(((13[0-9]{1})|(14[0-9]{1})|(15[0-9]{1})|(16[0-9]{1})|(18[0-9]{1}))+[0-9]{8})$/.test(num);
                },
                isUrl: function(str) {
                    return /([\w-]+\.)+[\w-]+.([^a-z])(\/[\w-.\/?%&=]*)?|[a-zA-Z0-9\-\.][\w-]+.([^a-z])(\/[\w-.\/?%&=]*)?/i.test(str);
                },
                isEmail: function(str) {
                    return /^([a-zA-Z0-9_\.\-\+])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})+$/.test(str);
                },
                isUsername: function(str) {
                    return /^[a-zA-Z6]{6,20}$/.test(str);
                },
                minLength: function(str, length) {
                    var strLength = $.getStrLength(str);
                    return strLength >= length;
                },
                maxLength: function(str, length) {
                    var strLength = $.getStrLength(str);
                    return strLength <= length;
                },
                trim: function(text){
                    return typeof(text) == "string" ? text.replace(/^\s*|\s*$/g, "") : text;
                },
                loading: function(w){
                    w = typeof(w) == 'undefined' ? 16 : w;
                    return '<img src="'+zrb.static_path+'/images/loading.gif" style="width:'+w+'px;" />';
                },
                //窗口事件
                fixEvent: function(e){
                    var evt = (typeof e == "undefined") ? window.event : e;
                    return evt;
                },
                //事件链接
                srcElement:	function(e){
                    if (typeof e == "undefined") e = window.event;
                    var src = document.all ? e.srcElement : e.target;
                    return src;
                },
                // Cookie相关
                setCookie: function(name, value, time) {
                    var exp = new Date();
                    var count = 60000;
                    exp.setTime(exp.getTime() + parseInt(count) * time);
                    document.cookie = name + "=" + value + "; expires=" + exp.toGMTString() + "; path=/";
                },
                getCookieVal: function(offset) {
                    var endstr = document.cookie.indexOf(";", offset);
                    if (endstr == -1) endstr = document.cookie.length;
                    return unescape(document.cookie.substring(offset, endstr));
                },
                delCookie: function(name) {
                    var exp = new Date();
                    exp.setTime(exp.getTime() - 864500000);
                    $.setCookie(name, "");
                },
                getCookie: function(name) {
                    var arg = name + "=";
                    var alen = arg.length;
                    var clen = document.cookie.length;
                    var i = 0;
                    while (i < clen) {
                        var j = i + alen;
                        if (document.cookie.substring(i, j) == arg) {
                            return $.com.util.getCookieVal(j);
                        }
                        i = document.cookie.indexOf(" ", i) + 1;
                        if (i == 0) {
                            break;
                        }
                    }
                    return null;
                },
                timejump: function(url , time){
                    if(typeof(time) != undefined){
                        setTimeout(function(){
                            url == '-1' ? parent.location.reload() : $.com.util.redirect(url);
                        } , time*1000);
                        return false;
                    }
                    parent.location.reload();
                    //url == '-1' ? parent.location.reload() : $.com.util.redirect(url);
                },
                redirect: function(uri, toiframe) {
                    if(toiframe != undefined){
                        $('#' + toiframe).attr('src', uri);
                        return false;
                    }
                    parent.location.href = uri;
                }
            }
        });
    $.com.common = {
        settings: {
            sms_code: ".J_sms_code",
            email_code: ".J_email_code",
            ajax_form: ".J_ajax_form",
            ajax_dialog: ".J_dialog",
            ajax_confirm : '.J_ajax_confirm',
            check_all: '.J_check_all',
            check_item: '.J_check_item',
            layer_tips: '.J_tips'
        },
        init: function(){
            $.com.common.init_input();
            $.com.common.ajax_dialog();
            $.com.common.ajax_submit_form();
            $.com.common.ajax_confirm();
        },
        init_input: function(){
            $(document).on("keyup", "input", function(){
                if(typeof($(this).attr("data-filter")) != 'undefined'){
                    var val = $(this).val();
                    switch($(this).attr("data-filter")){
                        case 'int':
                            var reg = /[^\d]+/;
                            val = val.replace(reg , '');
                            val = val ? parseInt(val) : 0;
                            break;
                        case 'float':
                            var reg = /[^\d\.]+/;
                            val = val.replace(reg , '');
                            val = val.match(/\d+(\.)?(\d+)?/g);
                            //val = val ? parseFloat(val) : 0;
                            break;
                        case 'chinese':
                            var reg = /[^\u4e00-\u9fa5]+/;
                            val = val.replace(reg , '');
                            break;
                        default:
                            break;

                    }
                    $(this).val(val);
                }
            });
            $(document).on('click','.J_replace',function(){
                window.location.reload();
            })
        },
        ajax_submit_form: function(){
            var J = $.com.common.settings;
            $(document).on("submit", J.ajax_form, function(){
                var _this_form = this,
                    _pos = $(_this_form).attr("pos"),
                    _pos = typeof(_pos)=='undefined' ? 2 : _pos,
                    _clear = $(_this_form).attr("clear"),
                    _cansubmit = true;
                layer.closeAll();
                $(_this_form).find("input[type='text'],input[type='number'],input[type='password'],select,textarea").each(function(){
                    var require = $(this).attr("require");
                    var thisvalue = $(this).val();
                    if(typeof(require)!='undefined' && require != '' && thisvalue == ''){
                        layer.msg(require);
                        _cansubmit = false;
                        return false;
                    }
                });
                if ($(_this_form).hasClass('disabled')) {
                    return false;
                }
                $(_this_form).addClass('disabled');
                var btnTxt = $(_this_form).find("input[type='submit'],button[type='submit']").val();
                if(_cansubmit){
                    $(_this_form).ajaxSubmit({
                        dataType: 'json',
                        resetForm: false,
                        cache: false,
                        type:'post',
                        data:{ajax:1, r:Math.random()},
                        timeout:15000, // 15s
                        beforeSubmit:function(){
                            $(_this_form).find("input[type='submit'],button[type='submit']")
                                .attr("disabled" , true)
                                .addClass("disabled")
                                .val("提交中...");
                            //layer.msg('正在提交中', {icon:16,time:30000});
                        },
                        success:function(res){
                            if(res.status==1){
                                if(res.url == ''){
                                    layer.msg(res.info);
                                }else if(res.url == '-1'){
                                    layer.msg(res.info, {time:1500},function(){
                                        parent.location.reload();
                                    });
                                } else {
                                    layer.msg(res.info, {time:1500},function(){
                                        window.location.href = res.url;
                                    });
                                }
                            }else{
                                layer.msg(res.info, {time:2000});
                            }
                            $(_this_form).find("input[type='submit'],button[type='submit']")
                                .removeAttr("disabled")
                                .removeClass("disabled")
                                .val(btnTxt);
                            $(_this_form).removeClass('disabled');
                        },
                        error: function(a , b , c){
                            if(b == 'timeout'){
                                layer.msg('超时了，网络偶有抽风！', {icon:16,time:2000});
                            }
                        }
                    });
                }
                return false;
            })
        },
        ajax_dialog: function(){
            var J = $.com.common.settings;
            $(document).on("click", J.ajax_dialog, function(){
                var _this = this;
                if($(_this).hasClass("disabled")){
                    return false;
                }
                $(_this).addClass("disabled");
                var _title = $(_this).attr("_title"),
                    _title = typeof(_title) != 'undefined' ? _title : $(_this).text(),
                    _url = $(_this).attr("_url"),
                    _url = typeof(_url)=='undeifned' ? '' : _url,
                    _width = $(_this).attr("_width"),
                    _width = typeof(_width) != 'undefined' ? _width : 500,
                    _height = $(_this).attr("_height"),
                    _height = typeof(_height) != 'undefined' ? _height : 300,
                    _batch = $(_this).attr("_batch"),
                    _batch = typeof(_batch) != 'undefined' ? true : false,
                    _full = $(_this).attr("_full"),
                    _full = typeof(_full) != 'undefined' ? true : false,
                    _html = $(_this).attr("_html"),
                    _html = typeof(_html) != 'undefined' ? true : false,
                    _noshade = $(_this).attr("_noshade"),
                    _noshade = typeof(_noshade) == 'undefined' ? 0.5 : false;
                _sclose = $(_this).attr("_sclose"),
                    _sclose = typeof(_sclose) != 'undefined' ? true : false;
                if(_url){
                    var _ids = '';
                    if(_batch){
                        _ids = $.com.common.check_ids();
                        if(_ids == ''){
                            layer.msg("请至少选择一项数据！");
                            $(_this).removeClass("disabled");
                            return false;
                        }
                        _url = _url+'?dialog=1&ids='+_ids;
                    }else{
                        _url = _url+'?dialog=1';
                    }
                    if(_html){
                        $.getJSON(_url, {ajax:1}, function(res){
                            $(_this).removeClass("disabled");
                            if(res.status == 1){
                                var dialogIndex = layer.open({
                                    type: 1,
                                    title: _title,
                                    //skin: 'layui-layer-rim',
                                    area: [_width+'px', _height+'px'],
                                    content: res.data,
                                    success: function(layero, index){
                                        $(layero).find(".J_position").trigger("click");
                                    }
                                });
                            }else{
                                layer.msg(res.info);
                            }
                        })
                    }else{
                        $(_this).removeClass("disabled");
                        var dialogIndex = layer.open({
                            type: 2,
                            title: _title,
                            shade: _noshade,
                            shadeClose: _sclose,
                            //skin: 'layui-layer-rim',
                            area: [_width+'px', _height+'px'],
                            content: _url,
                            maxmin: _full
                        });
                    }
                    _full && layer.full(dialogIndex);
                }else{
                    layer.msg("请求地址为空！");
                }
                return false;
            })
        },
        // ajax操作
        ajax_confirm: function(){
            var J = $.com.common.settings;
            $(document).on("click", J.ajax_confirm, function(){
                var _this = this,
                    url = $(_this).data("url"),
                    tip = $(_this).data("tip");
                layer.confirm(tip, {title:'确认提示'}, function(index){
                    $.ajax({
                        data: {},
                        type: "get",
                        url: url,
                        dataType: "json",
                        timeout: 15000,//15s
                        beforeSend: function(){
                            layer.msg('正在提交中~', {time:30000});
                        },
                        success: function(res){
                            if(res.status==1){
                                layer.msg(res.info, {time:1500}, function(){
                                    window.location.reload();
                                });
                            }else{
                                layer.msg(res.info, {time:5000});
                            }
                        },
                        error: function(a, b, c){
                            if(b == 'timeout'){
                                layer.msg('抱歉，超时了，网络偶有抽风', {time:2000});
                            }
                        }
                    })
                });
                return false;
            })
        },
        // date_picker: function(){
        //     $(document).on("focus", ".J_year", function(){
        //         WdatePicker({
        //             dateFmt:'yyyy'
        //         });
        //     });
        //     $(document).on("focus", ".J_month", function(){
        //         WdatePicker({
        //             dateFmt:'yyyy-MM'
        //         });
        //     });
        //     $(document).on("focus", ".J_day", function(){
        //         WdatePicker({});
        //     });
        //     $(document).on("focus", ".J_hour", function(){
        //         WdatePicker({
        //             dateFmt:'yyyy-MM-dd HH'
        //         });
        //     });
        //     $(document).on("focus", ".J_minute", function(){
        //         WdatePicker({
        //             dateFmt:'yyyy-MM-dd HH:mm'
        //         });
        //     });
        //     $(document).on("focus", ".J_second", function(){
        //         WdatePicker({
        //             dateFmt:'yyyy-MM-dd HH:mm:ss'
        //         });
        //     });
        // }
    };
    $.com.common.init();
})(jQuery);