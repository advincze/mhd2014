var Score = (function () {
	var current;
	return {

		init: function () {
			console.log("Score initialized");
			current = 0;
		},
		increment: function(type) {
			// type
			// 0: falsch -> 0 punkte
			// 1: richtig mit joker -> 5
			// 2: richtig ohne joker -> 10
			switch (type) {
				case 0: break;
				case 1: current += 5;break;
				case 2: current +=10;break;
				default:break;
			}
		},
		getScore: function() {
			return current;
		}
	};
	
})();
Score.init();