var maxCount    = 4
var minCount    = 1
var question    = "";
var time        = 0;
var displayed   = 0; //False
var clicked     = 0; //False
var counterTime = 6;
var sucFn       = "";
var startTime   = 0;

function initialize (ques, fn) {
   $("#loading").css('display','block');
   // clear the clock to 0s 
   question  = ques;
   sucFn     = fn; // Function to be called on success
   time      = 0;
   displayed = 0; //False
   clicked   = 0; //False
   startTime = new Date ().getTime ()
}

function removeCopy () {
	$("textarea[name='speedInput']").bind('copy paste', function (e) {
		e.preventDefault();
	 });
}

function showTime() {
        $("label[name=clock]").text(time + " s");
}

function validate() {

    // Place validation Functions here
    var compareString = $("textarea[name='speedInput']").val();
    if (question.indexOf(compareString)==-1) {
        $("label[name='question']").css('color', 'red');
    } else  {
        $("label[name='question']").css('color', 'green');
    }

    if(question.toString() == compareString.toString()) {	
	if (displayed == 0) {
             sucFn();	
	}
    }
}

function tick() {
	showTime();
	validate();
        currTime = new Date().getTime()
        time     = currTime - startTime
        time     = time/1000 
	startCounting();
}

function startCounting() {
    if (displayed == 0) {
 	setTimeout(tick,1);
    }
}

