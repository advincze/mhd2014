function Questions() {
    this.init = function() {
        $('.question-wrapper').each(function() {
            //TODO: initialize each question 
        })
    },
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
        console.log(obj);
    }
}
