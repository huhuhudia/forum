import showdown from 'showdown'

export const dateFormate =  function (t) {
  let date = new Date(t * 1000);

  let datestring = date.toLocaleDateString("zh-CN");
  let timestring = date.toLocaleTimeString("zh-CN");
  return  datestring + " " + timestring;
}

export const  md2Html=  function (src){

    let converter = new showdown.Converter();

    let html = converter.makeHtml(src);
    return  html;
}




