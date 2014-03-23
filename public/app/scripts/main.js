var questions,
    results = [],
    currentQuestion = 0,
    progresstimer;
$(document).ready(function() {
    questions = new Questions();
    questions.init();
    initEvents();
    initSizes();
});
$(window).load(function () {
    $('#loader').hide();
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
        clearTimeout(progresstimer);
        console.log("clrea");
            handleQuestionAnswer(correctAnswer, 600);
       
    });
    $('#start-btn').on('click', function() {
        startApp();
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
    console.log("prog");
    $(".question-wrapper.active .progress-bar.counter").removeClass("counter");
        progresstimer = setTimeout(function() {
            console.log("timeout");
                        handleQuestionAnswer(false,1);
        },15000);
}

function showEndScreen(load) {
    //$('.questions').css('visibility', 'hidden');
    $('.endscreen').show();
    console.log('ENDE!');
}