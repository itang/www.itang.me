$(function(){
  $.get("/todo_count",function(resp){
    $('#ic').html("(" + resp + ")");
  });
});
