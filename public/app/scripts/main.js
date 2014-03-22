$(document).ready(function() {
    var questions = new Questions();
    questions.init();
    questions.activate();
    initEvents();
    $('.answers-images .answer').css('height', parseInt($('.question-wrapper').css('width')) / 2);
    $(".progress-bar.counter").removeClass("counter").on("transitionend webkitTransitionEnd", function() {
        //alert("zeit abgelaufen");
        $(".progress-bar.counter").removeClass("counter").on("transitionend webkitTransitionEnd", function() {
            //alert("zeit abgelaufen");
            $(this).addClass("finished");
        });
    });
});

function initEvents() {
    $(window).on('resize', function() {
        $('.answers-images .answer').css('height', parseInt($('.question-wrapper').css('width')) / 2);
    });
    $('#all-questions').on('click', '.answer', function() {
        var questions = new Questions();
        questions.validate($(this));
    });
}
