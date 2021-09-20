/*
* MONO - A premium template from Designova
* Author: Designova, http://www.designova.net
* Copyright (C) 2017 Designova
* This is a premium product. For licensing queries please contact info@designova.net
*/

/*global $:false */
/*global window: false */
(function() {
    "use strict";
    $(function($) {

        //Detecting viewpot dimension
        var vH = $(window).height();
        var vW = $(window).width();
        //Adjusting Intro Components Spacing based on detected screen resolution
        $('.fullwidth').css('width', vW);
        $('.fullheight').css('height', vH);
        $('.halfwidth').css('width', vW / 2);
        $('.halfheight').css('height', vH / 2);

		// Page ready
		$(document).ready(function()
		{
			$('body').append('<div id="page-loading-polydoms-notifaction"></div>'); // Add page loading UI
			$('.polydom-fill-screen').css('height', $(window).height()+'px'); // Set initial hero height
			$('#scroll-hero').on("click", function(event)
			{
				event.preventDefault();
				$('html,body').animate({scrollTop: $('#scroll-hero').closest('.polydom').height()}, 'slow');
			});
			
			extraNavFuncs(); // Extra Nav Functions
			setUpSpecialNavs(); // Set Up Special NavBars 
			setUpDropdownSubs(); // Set Up Dropdown Menu Support
			setUpLightBox(); // Add lightbox Support
			setUpVisibilityToggle(); // Add visibility Toggle Support
			addSwipeSupport(); // Add Swipe Support	
			addKeyBoardSupport(); // Add Keyboard Support - Used for Lightbox Nav


		    //PRELOADER
		    $('body, html').addClass('preloader-running');
		    $('#mastwrap').css('visibility', 'hidden');
		    $(window).on("load", function() {
		            $("#status").fadeOut();
		            $("#preloader").delay(1000).fadeOut(1000);
		            $('body, html').removeClass('preloader-running');
		            $('body, html').addClass('preloader-done');
		            $("#mastwrap").delay(1000).css('visibility',
		                'visible');
		    });
		    $('.menu-holder a').on('click',function(){
		            $("#status").show();
		            $("#preloader").show();
		            $('body, html').addClass('preloader-running');
		    });
		});

		// Loading page complete
		$(window).on("load", function()
		{
			setFillScreenpolydomHeight();
			animateWhenVisible();  // Activate animation when visible	
			$('#page-loading-polydoms-notifaction').remove(); // Remove page loading UI
		}
		).on("resize", function() // Window resize 
		{		
			setFillScreenpolydomHeight();	
		}); 


    });
    // $(function ($)  : ends
})();
//  JSHint wrapper $(function ($)  : ends