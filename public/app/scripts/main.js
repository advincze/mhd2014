var questions;
$(document).ready(function() {
    questions = new Questions();
    questions.init();
    initEvents();
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

function startProgress() {
    $(".question-wrapper.active .progress-bar.counter").removeClass("counter").on("transitionend webkitTransitionEnd", function() {
        $(this).addClass("finished");
    });
}

function showEndScreen(load) {
    if (load) {
        //$.load ...
    } else {
        //hide other divs
    }
}

function showStartScreen(load) {

}
