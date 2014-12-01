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
        // result = '<img src="' + url + '"></img>';
        result = '该选项暂不支持，请看右侧说明。';
        // $('#result').html(result + "<br> 链接：" + url);
        $('#result').html(result);
        break;
      case 'svg':
        result = '<img src="'+ url + '">';
        $('#result').html(result + "<br> svg 资源链接：" + url);
        break;
      case 'mml':
        $.get(url, function(data){
          result = data;
          $('#result').html(result + "<br> mml 链接：" + url);
        });
        break;
      default:
        result = '输出类型错误！';
        $('#result').html(result + "<br> 链接：" + url);
    }
  })
});
