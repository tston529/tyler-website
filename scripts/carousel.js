var numStories=0;
$(document).ready(function (){
    $(".textBlock").each(function(){
        if(numStories > 0) {
            $(this).hide();
        }
        $(this).attr('id', numStories);
        numStories++;
    });
})

$('#next').click(function(){
    var found = 0;
    $(".textBlock").each(function(){
        if (found == 1) {
            found++;
            $(this).addClass("current");
            $(this).show("fast");
        }
        if (found == 0 && ($(this).attr('id') < numStories-1) && $(this).attr('class') == "textBlock current") {
            $(this).removeClass("current");
            $(this).hide("fast");
            found++;
        }
    })
})

$('#previous').click(function () {
    var found = 0;
    $($(".textBlock").get().reverse()).each(function () {
        if (found == 1) {
            found++;
            $(this).addClass("current");
            $(this).show("fast");
        }
        if (found == 0 && ($(this).attr('id') > 0) && $(this).attr('class') == "textBlock current") {
            $(this).removeClass("current");
            $(this).hide("fast");
            found++;
        }
    })
})