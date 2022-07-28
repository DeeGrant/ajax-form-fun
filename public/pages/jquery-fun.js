$(function () {
    $('#btn').click(function () {
        console.log($('#test').text())
        console.log($('#test').html())
        console.log($('#fcc').attr('href'))
    });

    $('#btn_set').click(function () {
        $('#test').text("freeCodeCamp is <b>awesome</b>!") // html doesn't work
    });

    $('#btn_set_2').click(function () {
       $('#test_2').html("freeCodeCamp is <b>Awesome</b>!") // html does work!
    });

    $('#btn_value').click(function () {
        $('#name').val("Dee Grant")
    })

    $('#btn_multi').click(function () {
        $('section p').text(function (index, original_val) {
            return index + ' - ' + original_val
        })
    })
});