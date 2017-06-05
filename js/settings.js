$(document).ready(function() {
    $('.cursor').click(function(event) {
        clickEvent(event);
    });
});

function saveApi() {
    var apiName = $('#addNewApiDiv #apiName').val();
    var beginningEscape = $('#addNewApiDiv #beginningEscape').val();
    var endingEscape = $('#addNewApiDiv #endingEscape').val();
    console.log(apiName);
    console.log(beginningEscape);
    console.log(endingEscape);
    $.ajax({
        url: "/updateSettings",
        type: "post",
        data: {
            apiName: apiName,
            beginningEscape: beginningEscape,
            endingEscape: endingEscape
        }
    });
}

function clickEvent(event) {
    var target = $(event.target);
    var targetId = target.attr('id');

    if(target.data('type') == "pre") {
        var type = "textarea";
    } else {
        var type = "input";
    }

    target.replaceWith($('<'+type+'/>', {
        'id': targetId,
        'class': target.attr('class'),
        'type': 'text',
        'value': target.text().trim(),
        'val': target.text().trim(),
        'data-type': target.data('type'),
        'data-original': target.text().trim()
    }));

    var target = $('#'+targetId);
    target.focus();
    target.blur(function(event) {
        blurEvent(event);
    });
}

function blurEvent(event) {
    var target = $(event.target);
    var targetId = target.attr('id');

    var confirmResult = window.confirm("Save changes?");

    target.replaceWith($('<'+target.data('type')+'/>', {
        'id': targetId,
        'class': target.attr('class'),
        'text': ((confirmResult) ? target.val() : target.data('original')),
        'data-type': target.data('type')
    }));
    var target = $('#'+targetId);
    target.click(function(event) {
        clickEvent(event);
    });
}
