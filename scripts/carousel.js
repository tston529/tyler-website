var numStories=0;
$(document).ready(function () {
    $(".textBlock").each(function(){
        if(numStories > 0) {
            $(this).hide();
        }
        $(this).attr('id', numStories);
        numStories++;
    });
    $("#previous").hide();
})

$('#next').click(function() {
    $("#previous").show("fast");
    var found = 0;
    $(".textBlock").each(function() {
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
    //console.log($(".current").first().attr('id') + " " + $(".textBlock").last().attr('id'));
    if ($(".current").first().attr('id') === $(".textBlock").last().attr('id')) {
        $("#next").hide("fast");
    }
})

$('#previous').click(function () {
    $("#next").show("fast");
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
    //console.log($(this).attr('id') + " " + $(".textBlock").first().attr('id'));
    if ($(".current").first().attr('id') === $(".textBlock").first().attr('id')) {
        $("#previous").hide("fast");
    }
})