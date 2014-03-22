$(document).ready(function() {
    var questions = new Questions();
    questions.init();
    //questions.activate();
    initEvents();
    $(".progress-bar.counter").removeClass("counter").on("transitionend webkitTransitionEnd", function() {
        //alert("zeit abgelaufen");

        $(this).addClass("finished");
        //$(".question-wrapper.active .answer").css("opacity", 0.6);
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
        }, 600);
    });
}

function initSizes() {
    $('.answers-images .answer').css('height', parseInt($('.question-wrapper').css('width')) / 2);
}
