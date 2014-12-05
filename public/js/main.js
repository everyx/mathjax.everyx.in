$(document).ready(function() {
  $('#go-button').on('click', function(e) {
    e.preventDefault();

    var expressionType = $('#expression-type').val();
    var outType = $('#out-type').val();
    var expression = btoa($('#expression').val());

    var url = document.URL + expressionType + '/' + outType + '/' + expression;

    var result;
    switch(outType) {
      case 'png':
        result = '<img src="' + url + '"></img>';
        $('#result').html(result + "<br> png 外链地址：" + url);
        break;
      case 'svg':
        result = '<object data="'+ url + '" type="image/svg+xml"></object>';
        $('#result').html(result + "<br> svg 外链地址：" + url);
        break;
      case 'mml':
        $.get(url, function(data){
          result = data;
          $('#result').html(result + "<br> mml 内容获取地址：" + url);
        });
        break;
      default:
        result = '输出类型错误！';
        $('#result').html(result + "<br> 链接：" + url);
    }
  })
});
