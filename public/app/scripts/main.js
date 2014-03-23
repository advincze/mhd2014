var questions,
    results = [],
    currentQuestion = 0;
$(document).ready(function() {
    questions = new Questions();
    questions.init();
    initEvents();
    initSizes();
});

function startApp() {
    $('.startscreen').hide();
<<<<<<< HEAD
=======
    $('.questions').css('visibility', '');
                        startProgress();

>>>>>>> be6ea1e529074a2637645f70c7111afd7ad95d4a
}

function initEvents() {
    $(window).on('resize', function() {
        $('.answers-images .answer').css('height', parseInt($('.question-wrapper').css('width')) / 2);
    });
    $('#all-questions').on('click', '.answer', function() {
        var correctAnswer = questions.validate($(this));
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
<<<<<<< HEAD
=======

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
>>>>>>> be6ea1e529074a2637645f70c7111afd7ad95d4a
