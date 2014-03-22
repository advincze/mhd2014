var questions;
$(document).ready(function() {
    questions = new Questions();
    questions.init();
    initEvents();
    console.log($(".question-wrapper.active .progress-bar.counter").hasClass("counter"));
    $(".question-wrapper.active .progress-bar.counter").removeClass("counter").on("transitionend webkitTransitionEnd", function() {
        //alert("zeit abgelaufen");

        $(this).addClass("finished");
        //questions.next();
        //$(".question-wrapper.active .answer").css("opacity", 0.6);
    });
    initSizes();
});

function initEvents() {
    $(window).on('resize', function() {
        $('.answers-images .answer').css('height', parseInt($('.question-wrapper').css('width')) / 2);
    });
    $('#all-questions').on('click', '.answer', function() {
        questions.validate($(this));
        window.setTimeout(function() {
            questions.next();
        }, 600);
    });
}

function initSizes() {
    $('.answers-images .answer').css('height', parseInt($('.question-wrapper').css('width')) / 2);
}
