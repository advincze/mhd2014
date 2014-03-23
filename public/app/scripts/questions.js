function Questions() {
    var currentQuestion = 0,
        questCount = 0,
        questMax = 4,
        questions = {};

    this.init = function() {
        $.ajax({
            //url: 'testData.json',
            url: "/data/daily",
            type: 'GET',
            dataType: 'json',
            success: function(data) {
                questions = data;
                questCount = questions.length;
                var first = $('#all-questions .question-wrapper').first();
                console.log(first);
                first.children('.question-text').children('h3').text(questions[0].headline);
                console.log(questions[0].answerRight);
                console.log(questions[0].answerWrong);
                first.children('.answers-images').children(".answer.right").attr('data-img', questions[0].rightImageURL.replace("w1-h1", "w500-h500-oo"));
                first.children('.answers-images').children(".answer.wrong").attr('data-img', questions[0].wrongImageURL.replace("w1-h1", "w500-h500-oo"));
                first.css('left', 0);
                var l = 100;
                console.log(questions.length);
                for (var i = 1; i < questions.length; i++) {
                    var clone = first.clone();
                    clone.children('.question-text').children('h3').text(questions[i].headline);
                    clone.children('.answers-images').children(".answer.right").attr('data-img', questions[i].rightImageURL.replace("w1-h1", "w500-h500-oo"));
                    clone.children('.answers-images').children(".answer.wrong").attr('data-img', questions[i].wrongImageURL.replace("w1-h1", "w500-h500-oo"));
                    clone.css('left', l + '%');
                    l += 100;
                    $('#all-questions .questions-inner-wrapper').append(clone);
                }
                first.addClass('active');
                //takes the active question and activates it
                $('.question-wrapper .answer').each(function() {
                    console.log($(this).attr('data-img'));
                    $(this).css('background-image', "url(" + $(this).attr('data-img') + ")");
                });

                
                

            }

        });
    },
    this.next = function() {

        var isEnd = true;

        if (currentQuestion < questMax) {
            isEnd = false;
            currentQuestion++;
            $('.question-wrapper.active').removeClass('active');
            var oldValue = $('.questions-inner-wrapper')[0].style.left == "" ? 0 : $('.questions-inner-wrapper')[0].style.left;
            var newValue = (parseInt(oldValue) - 100) + '%';
            $('.questions-inner-wrapper').css('left', newValue);
            $('.question-wrapper').eq(currentQuestion).addClass('active');
            console.log(currentQuestion);
            $('.badge-task').text(currentQuestion + 1 + " / 5");
             setTimeout(function () {
               startProgress();
            },500);
        }

        return isEnd;
    }
    this.getCurrent = function() {
        return currentQuestion;
    }
    this.validate = function(obj) {

        var correctAnswer = false;

        if (obj.hasClass('right')) {
            correctAnswer = true;
            score.increment(2);
            score.setScore();
            //obj.css('background', '#0F0');
            obj.append('<div class="overlay-right"></div>');
        } else {
            //obj.css('background', '#F00');
            obj.append('<div class="overlay-wrong"></div>');
        }

        return correctAnswer;
    }
    this.getQuestionData = function() {
        return questions;
    }
}
