function Questions() {
    this.init = function() {
        $.ajax({
            //url: 'testData.json',
            url: "/data/daily",
            type: 'GET',
            dataType: 'json',
            success: function(data) {
                console.log(data);

                var questions = data;
                var first = $('#all-questions .question-wrapper').first();
                var clone = first.clone();
                console.log(first);
                first.children('.question-text').children('h3').text(questions[0].headline);
                console.log(questions[0].answerRight);
                console.log(questions[0].answerWrong);

                first.children('.answers-images').children(".answer.right").attr('data-img', questions[0].rightImageURL.replace("w1-h1","w500-h500-oo"));
                first.children('.answers-images').children(".answer.wrong").attr('data-img', questions[0].wrongImageURL.replace("w1-h1","w500-h500-oo"));
                first.css('left', 0);
                var l = 100;
                console.log(questions.length);
                for (var i = 1; i < questions.length; i++) {
                    clone.children('.question-text').children('h3').text(questions[i].headline);
                    clone.children('.answers-images').children(".answer.right").attr('data-img', questions[i].rightImageURL.replace("w1-h1","w500-h500-oo"));
                    clone.children('.answers-images').children(".answer.wrong").attr('data-img', questions[i].wrongImageURL.replace("w1-h1","w500-h500-oo"));
                    clone.css('left', l + '%');
                    l += 100;
                    $('#all-questions .questions-inner-wrapper').append(clone);
                }
                //takes the active question and activates it
                $('.question-wrapper.active .answer').each(function() {
                    console.log($(this).attr('data-img'));
                    $(this).css('background-image', "url(" + $(this).attr('data-img') + ")");
                });
            }
        });
    },

    this.next = function() {
        //TODO: swipes to the next question
        var oldValue = $('.questions-inner-wrapper')[0].style.left == "" ? 0 : $('.questions-inner-wrapper')[0].style.left;
        var newValue = (parseInt(oldValue) - 100) + '%';
        $('.questions-inner-wrapper').css('left', newValue);
    }
    this.validate = function(obj) {
        if (obj.hasClass('right')) {
            //obj.css('background', '#0F0');
            obj.append('<div class="overlay-right"></div>');
        } else {
            //obj.css('background', '#F00');
            obj.append('<div class="overlay-wrong"></div>');
        }
    }
}
