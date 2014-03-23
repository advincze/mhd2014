var questions,
    results = [],
    progresstimer;
       var ajax = false,
        started = false;
$(document).ready(function() {
    questions = new Questions();
    questions.init();
    
    initEvents();
    initSizes();
});



function startApp() {
    $('.startscreen').hide();
    $('.questions').css('visibility', '');
    if(started && ajax) {
        startProgress();
    }
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
    $('.topic').on('click', function() {
        started=true;
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
function getHint() {
    console.log(questions.getQuestionData()[questions.getCurrent()].fullArticleURL);
    $.ajax({
        url:"http://api.embed.ly/1/extract?key=03748ef27f6d41639a1973fd03ebabbb&url="+questions.getQuestionData()[questions.getCurrent()].fullArticleURL+"&maxwidth=300&maxheight=300&format=json&callback=setHint"
    })
}
function setHint(data) {
    var hint = false;
    var crops = [];
    crops[questions.getQuestionData()[questions.getCurrent()].rightImageURL.match(/crop[0-9]*/)[0]] = questions.getQuestionData()[questions.getCurrent()].rightImageURL.match(/crop[0-9]*/)[0];
    for(var i = 0; i<data.images.length;i++) {
        questions.getQuestionData()[questions.getCurrent()].rightImageURL
        if (typeof crops[data.images[i].url.match(/crop[0-9]*/)[0]] != "string") {
            hint=true; // HINT VORHANDEN -> Button anzeigen
            console.log(data.images[i].url,"gnicht efunden");
            crops[data.images[i].url.match(/crop[0-9]*/)[0]] = data.images[i].url.match(/crop[0-9]*/)[0];
        } else {
        console.log(data.images[i].url,"gefunden");
        }
        // TODO FÜR MORITZ IM OVERLAY. ggf 1. Bild raus wegen zusätzlichem Crop des Teasers oder nach IDs in der URL wegen Dopplungen
    }
    
}
function stopTimer () {

}
function startProgress() {
    getHint();
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