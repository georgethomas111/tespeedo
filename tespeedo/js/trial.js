var Game = function () {

	// Register the handlers	
        var clockName  = 'result';
        var counterId  = 'counter';
	var startName  = 'start';

	this.killId    = -1;
	this.time      = 0;
	this.displayed = 0;

	}
	var StartGame  = function () {
	  alert ("called");
	  tick();
	 }
        var myStart    = document.getElementById (startName);
	if (typeof (myStart.addEventListener) === 'function') {
	     myStart.addEventListener ('click', StartGame);
	} else {
	     console.log ("failed to add event");
	}


	var showTime = function () {
		$("label[name=clock]").text(this.time + " s");
	}

	var validate = function () {

		// Place validation Functions here
		var compareString = $("textarea[name='speedInput']").val();
		if (question.indexOf(compareString)==-1) {
			$("label[name='question']").css('color', 'red');
		} else  {
			$("label[name='question']").css('color', 'green');
		}

		if(question.toString() == compareString.toString()) {	
				this.success();
			        finalyze ();	
		}
	}

	var tick = function () {
		showTime();
		validate();
		this.currTime = new Date().getTime()
		this.time     = this.currTime - startTime
		this.time     = this.time/1000 
		startCounting();
	}

	var startCounting = function () {
		if (displayed == 0) {
		      this.killId = window.setTimeout(tick,1);
		}
	}
	
	var finalyze = function () {
		window.clearTimeout(this.killId)
	}
} ();
