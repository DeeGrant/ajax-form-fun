function formSubmit() {
    const data = $('form')
        .serializeArray()
        .reduce((obj, element) => {
            obj[element.name] = element.value
            return obj
        }, {})

    $.post('http://localhost:8000/formSubmit',
        JSON.stringify(data),
        function (data, success) {
            data.forEach(thing => {
                $('#thing-list').append('<li>' + thing.name + ' - ' + thing.city + ' - ' + thing.things + '</li>')
            })
        })
}

$(function () {
    $('form button').click(formSubmit)
});