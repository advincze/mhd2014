function Questions() {
    this.init = function() {
        $.ajax({
            url: 'testData.json',
            type: 'GET',
            dataType: 'json',
            success: function(data) {
                //console.log(data.questions);
                var questions = data.questions;
                var first = $('#all-questions .question-wrapper').first();
                var clone = first.clone();
                first.children('.question-text').children('h3').text(questions[0].headline);
                first.children('.answers.right').attr('data-img', questions[0].answerRight);
                first.children('.answers.wrong').attr('data-img', questions[0].answerWrong);
                first.css('left', 0);
                var l = 100;
                for (var i = 1; i < questions.length; i++) {
                    clone.children('.question-text').children('h3').text(questions[i].headline);
                    clone.children('.answers.right').attr('data-img', questions[i].answerRight);
                    clone.children('.answers.wrong').attr('data-img', questions[i].answerWrong);
                    clone.css('left', l + '%');
                    l += 100;
                    $('#all-questions .questions-inner-wrapper').append(clone);
                }
            }
        });
    },
    this.activate = function() {
        //takes the active question and activates it
        $('.question-wrapper.active .answer').each(function() {
            $(this).css('background-image', "url(" + $(this).attr('data-img') + ")");
        })
    },
    this.next = function() {
        //TODO: swipes to the next question
        var oldValue = $('.questions-inner-wrapper')[0].style.left == "" ? 0 : $('.questions-inner-wrapper')[0].style.left;
        var newValue = (parseInt(oldValue) - 100) + '%';
        $('.questions-inner-wrapper').css('left', newValue);
    }
    this.validate = function(obj) {
        if (obj.hasClass('right')) {
            obj.css('background', '#0F0');
        } else {
            obj.css('background', '#F00');
        }
    }
}
