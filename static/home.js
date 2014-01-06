$( document ).ready(function() {
	$("#signupForm").submit(function(event){
		var address = $("#signupForm input").val()
		$.post("/new", {email: address})
			.done(function(data){
				$( ".result" ).html( data );
				$("#signupForm").hide();
			});
		event.preventDefault();
	});
});