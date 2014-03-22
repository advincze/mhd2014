function Questions() {
    this.init = function() {
        $.ajax({
            url: 'testData.json',
            type: 'GET',
            dataType: 'json',
            success: function(data) {
                console.log(data.questions);
                //this.initQuestions(data.questions);
            }
        });
    },
    this.initQuestions = function() {}
    this.activate = function() {
        //takes the active question and activates it
        $('.question-wrapper.active .answer').each(function() {
            $(this).css('background-image', "url(" + $(this).attr('data-img') + ")");
        })
    },
    this.next = function() {
        //TODO: swipes to the next question
    }
    this.validate = function(obj) {
        if (obj.hasClass('right')) {
            obj.css('background', '#0F0');
        } else {
            obj.css('background', '#F00');
        }
    }
}
