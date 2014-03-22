function Score () {
    var current;

    this.init = function() {
        current = 0;
        console.log("Score inizialized");
    }
    

    this.getScore = function() {
        return current;
    };
    this.setScore = function () {
        $(".badge-score").text(parseInt(current));
    }
    this.increment =  function(type) {
        // 0 falsch -> +0
        // 1 mit joker richtig -> +5
        // 2 richtig -> +10
        switch (type) {
            case 0:
                break;
            case 1:
                current += 5;
                break;
            case 2:
                current += 10;
                break;
            default:
                break;
        }
    }
};

var score = new Score();
score.init();