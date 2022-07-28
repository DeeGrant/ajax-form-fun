// Wait to run until the entire document has loaded
// deprecated syntax:
// $(document).ready(function () {
// });
// New form:
$(function () {
    $('p').html("Hello World").on('click mouseenter', function () {
        console.log('paragraph stuff')
    });

    // turn off mouse-enter handler from above
    $('p').off('mouseenter')

    let i = 0

    $('h1').hover(function () {
        console.log('what?', i++)
        $('h1').css('color', '#aaa')
    },
    function () {
        console.log('who?', i++)
        $('h1').css('color', '#000')
    });

    $('h2').on({
        'click': function () {
            console.log('Object click event happened')
        },
        'mouseover': function () {
            console.log('Object hover event happened')
        },
    })

    $('input').focus(function () {
        console.log('Something happened!')
    });

    $('#first').click(function () {
        $.get(
            'http://localhost:8000/getStuff', {
                name: "GFG"
            },
            function (data) {
                console.log(data)
                $('p').html(data.name)
            })
    })

    $('#second').click(function () {
        alert('POST!')
    })

    $('#once').one('click', function () {
        console.log('I only run this event once!')
    });
});
