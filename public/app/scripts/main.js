var questions,
    results = [],
    progresstimer;

$(document).ready(function() {
    questions = new Questions();
    questions.init();
    
    initEvents();
    initSizes();
});



function startApp() {
    $('.startscreen').hide();
    $('.questions').css('visibility', '');
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
    $(".question-wrapper.active .progress-bar.counter").removeClass("counter");
    progresstimer = setTimeout(function() {
        console.log("timeout");
        handleQuestionAnswer(false,1);
    },15000);
}

function showEndScreen() {
    //$('.questions').css('visibility', 'hidden');
    $('.endscreen').show();

    var rightAnswers = 0;
    results.forEach(function(e){
        if(e) rightAnswers++;
    });

    $('.endscore').text(rightAnswers);

    var $resultList = $('.endscreen').find('.results'),
        questionData = questions.getQuestionData();

    $.each($resultList.find('li'), function(i,e){  
        var data = questionData[i],
            element = $(e),
            $resultText = element.find('.result-text');
        element.find('.result-img').css('background-image', 'url(' + data.rightImageURL + ')');
        $resultText.find('.hl').text(data.headline);
        $resultText.find('.articleBtn').attr('href', data.fullArticleURL);
        
        if(results[i]) $(e).addClass('correct');
    });


    console.log('ENDE!');
}