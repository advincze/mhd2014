$(document).ready(function() {
    var questions = new Questions();
    questions.init();
    questions.activate();
    initEvents();
    $(".progress-bar.counter").removeClass("counter").on("transitionend webkitTransitionEnd", function() {
        //alert("zeit abgelaufen");
        $(".progress-bar.counter").removeClass("counter").on("transitionend webkitTransitionEnd", function() {
            //alert("zeit abgelaufen");
            $(this).addClass("finished");
        });
    });
    initSizes();
});

function initEvents() {
    $(window).on('resize', function() {
        $('.answers-images .answer').css('height', parseInt($('.question-wrapper').css('width')) / 2);
    });
    $('#all-questions').on('click', '.answer', function() {
        var questions = new Questions();
        questions.validate($(this));
        window.setTimeout(function() {
            questions.next();
        }, 2000);
    });
}

function initSizes() {
    $('.answers-images .answer').css('height', parseInt($('.question-wrapper').css('width')) / 2);
}
