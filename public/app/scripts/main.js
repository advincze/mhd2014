var questions,
    results = [],
    currentQuestion = 0;
$(document).ready(function() {
    showStartScreen(true);
    questions = new Questions();
    questions.init();
    initEvents();
    initSizes();
});

function startApp() {
    $('.startscreen').hide();
    $('.questions').css('visibility', '');
                        startProgress();

}

function initEvents() {
    $(window).on('resize', function() {
        $('.answers-images .answer').css('height', parseInt($('.question-wrapper').css('width')) / 2);
    });
    $('#all-questions').on('click', '.answer', function() {
        var correctAnswer = questions.validate($(this));
        handleQuestionAnswer(correctAnswer, 600);
       
    });
}

function initSizes() {
    $('.answers-images .answer').css('height', parseInt($('.question-wrapper').css('width')) / 2);
}

function handleQuestionAnswer(correctAnswer, timeout){
        results[questions.getCurrent()] = correctAnswer;
        
        window.setTimeout(function() {
            var isEnd = questions.next()
            if(isEnd){
                showEndScreen();
            }
        }, timeout);
}

function startProgress() {
    $(".question-wrapper.active .progress-bar.counter")
        .removeClass("counter")
        .on("transitionend webkitTransitionEnd", function() {
            $(this).addClass("finished");
            handleQuestionAnswer(false,1);
        });
}

function showEndScreen(load) {
    $('.questions').css('visibility', 'hidden');
    $('.endscreen').show();
    console.log('ENDE!');
}

function showStartScreen(load) {
    $('.questions').css('visibility', 'hidden');
    $('.endscreen').hide();
    if (load) {
        $('div.startscreen').load("templates/startscreen.html", function(response) {
            $('#start-btn').on('click', function() {
                startApp();
            });
        });
    } else {

    }
}
