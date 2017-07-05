$(document).ready(function() {
    $('.cursor').click(function(event) {
        clickEvent(event);
    });
});

function updateElement(target) {
    sendUpdateElementRequest(target.attr('id'), target.val(), false);
}

function sendUpdateElementRequest(id, value, closeModalNeeded) {
    // @TODO Consolidate code
    $.ajax({
        url: "/updateSettings",
        type: "post",
        data: {
            action: "updateField",
            id: id,
            value: value
        }
    }).done(function(response) {
        console.log(response);
        if(closeModalNeeded) {
            closeModal();
        }
    });
}

function saveApi() {
    var apiName = $('#addNewApiDiv #apiName').val();
    var beginningEscape = $('#addNewApiDiv #beginningEscape').val();
    var endingEscape = $('#addNewApiDiv #endingEscape').val();
    // @TODO Consolidate code
    $.ajax({
        url: "/updateSettings",
        type: "post",
        data: {
            action: "saveApi",
            apiName: apiName,
            beginningEscape: beginningEscape,
            endingEscape: endingEscape
        }
    }).done(function(response) {
        var json = $.parseJSON(response);
        if(json.Status) {
            closeModal();
        } else {
            alert("There was an error saving!");
        }
    });
}

function saveMessage(apiId) {
    var identifier = $('#identifier-'+apiId).val();
    $.ajax({
        url: "/updateSettings",
        type: "post",
        data: {
            action: "saveMessage",
            apiId: apiId,
            identifier: identifier
        }
    }).done(function(response) {
        var json = $.parseJSON(response);
        if(json.Status) {
            closeModal();
        } else {
            alert("There was an error saving!");
        }
    });
}

function saveResponse(messageId) {
    var response = $('#response-'+messageId).val();
    var isDefault = $('#default-'+messageId).val();
    var condition = $('#condition-'+messageId).val();
    $.ajax({
        url: "/updateSettings",
        type: "post",
        data: {
            action: "saveResponse",
            messageId: messageId,
            response: response,
            isDefault: isDefault,
            condition: condition
        }
    }).done(function(response) {
        var json = $.parseJSON(response);
        if(json.Status) {
            closeModal();
        } else {
            alert("There was an error saving!");
        }
    });
}

function clickEvent(event) {
    var target = $(event.target);
    var targetId = target.attr('id');

    if(target.data('modal')) {
        window.location = "#"+target.data('modal');
        window.location.reload();
        return;
    }

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
    var confirmResult = window.confirm("Save changes?");

    if(confirmResult) {
        updateElement(target);
    } else {
        var targetId = target.attr('id');
        // @TODO Move into function
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
}

// Save a provided field type from a modal
function saveModal(field, id) {
    var textArea = $('#'+field+id+"TextArea");
    var oldValue = $('#'+field+id+"OldValue");
    if(textArea.val() == oldValue.val()) {
        closeModal();
    } else {
        sendUpdateElementRequest(field+"Field"+id, textArea.val(), true);
    }
}

function closeModal() {
    window.location = "#close";
    window.location.reload();
}

function insertValueIntoTextArea(id, type, value) {
    var templateBox = document.getElementById(type+id+'TextArea');
    var cursor = templateBox.selectionStart;
    var cursorEnd = templateBox.selectionEnd;
    var prior = (templateBox.value).substring(0, cursor);
    var after = (templateBox.value).substring(cursorEnd, templateBox.value.length);
    templateBox.value = prior + value + after;
    closeModalModal();
    templateBox.selectionStart = templateBox.selectionEnd = cursor;
    templateBox.focus();
}

function displayFieldDialog(type, identifier) {
    $('#newFieldModal'+type+'-'+identifier).css('pointer-events', 'auto');
    $('#newFieldModal'+type+'-'+identifier).css('opacity', 1);
    $(document).keyup(function(e) {
        if(e.keyCode == 27) {
            closeModalModal();
        }
    });
}

function closeModalModal() {
    $('.modalModal').css('opacity', 0);
    $('.modalModal').css('pointer-events', 'none');
    // Disable escape key
    $(document).keyup(function(e) {
    });
}

function saveNewField(id, type, beginningEscape, endingEscape) {
    $.ajax({
        url: "/updateSettings",
        type: "post",
        data: {
            action: "saveNewField",
            type:  $('#fieldType'+type+'-'+id).val(),
            id: $('#fieldId'+type+'-'+id).val(),
            value: $('#newFieldInput'+type+'-'+id).val()
        }
    }).done(function(response) {
        responseObject = JSON.parse(response);
        insertValueIntoTextArea(
            id,
            type.toLowerCase(),
            beginningEscape + responseObject.Id + endingEscape
        );
    });
}
