$(document).ready(function() {
    var questions = new Questions();
    questions.init();
    questions.activate();
    initEvents();
    $('.answers-images .answer').css('height', parseInt($('.question-wrapper').css('width')) / 2);
});

function initEvents() {
    $(window).on('resize', function() {
        $('.answers-images .answer').css('height', parseInt($('.question-wrapper').css('width')) / 2);
    });
    $('.answer').on('click', function() {
        var questions = new Questions();
        questions.validate(this);
    });
}
