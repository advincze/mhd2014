var Score = function() {
    var current;

    function init() {
        current = 0;
        console.log("Score inizialized");
    }
    init();

    this.getScore = function() {
        return current;
    };
    increment: function(type) {
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
