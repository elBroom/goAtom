$(function() {
    checkToken();
    getHistory();
    setInterval(getHistory, 5000); // 5 sec

    $('#logout').click(function(e) {
        clearError();
        clearSuccess();
        $.ajax({
            url: pref+'/logout',
            method: 'POST',
            contentType: "application/json; charset=utf-8",
            dataType: 'json',
            data: JSON.stringify(''),
            beforeSend: function setAuthHeader(xhr){
                xhr.setRequestHeader('Authorization', readCookie("token"));
            },
            success: function(data, textStatus, jqXHR) {
                debugger;
            },
            error: function(jqXHR, textStatus, errorThrown) {
                if (jqXHR.responseText == "ok") {
                    eraseCookie("token");
                    window.location = login_link;
                } else {
                    viewError(jqXHR.responseText);
                }
            }
        });
        e.preventDefault(e);
    });

    $('#insert_data').click (function(e) {
        var data = getFormData($('#data-form'));
        var err = "";
        clearError();
        clearSuccess();

        if (data.key == "") {
            viewError("Key is required field");
            return false;
        } else if (data.value == "") {
            viewError("Value is required field");
            return false;
        }

        $.ajax({
            url: pref+'/value/',
            method: 'POST',
            contentType: "application/json; charset=utf-8",
            dataType: 'json',
            data: JSON.stringify(data),
            beforeSend: function setAuthHeader(xhr){
                xhr.setRequestHeader('Authorization', readCookie("token"));
            },
            success: function(data, textStatus, jqXHR) {
                viewSuccess("Insert with key "+data.key);
            },
            error: function(jqXHR, textStatus, errorThrown) {
                viewError(jqXHR.responseText);
            }
        });
        e.preventDefault(e);
    });

    $('#update_data').click (function(e) {
        var data = getFormData($('#data-form'));
        var err = "";
        clearError();
        clearSuccess();

        if (data.key == "") {
            viewError("Key is required field");
            return false;
        }  else if (data.value == "") {
            viewError("Value is required field");
            return false;
        }

        $.ajax({
            url: pref+'/value/'+data.key,
            method: 'PUT',
            contentType: "application/json; charset=utf-8",
            dataType: 'json',
            data: JSON.stringify(data),
            beforeSend: function setAuthHeader(xhr){
                xhr.setRequestHeader('Authorization', readCookie("token"));
            },
            success: function(data, textStatus, jqXHR) {
                viewSuccess("Update with key "+data.key);
            },
            error: function(jqXHR, textStatus, errorThrown) {
                viewError(jqXHR.responseText);
            }
        });
        e.preventDefault(e);
    });

    $('#get_data').click (function(e) {
        var data = getFormData($('#data-form'));
        var err = "";
        clearError();
        clearSuccess();

        if (data.key == "") {
            viewError("Key is required field");
            return false;
        }

        $.ajax({
            url: pref+'/value/'+data.key,
            method: 'GET',
            contentType: "application/json; charset=utf-8",
            dataType: 'json',
            beforeSend: function setAuthHeader(xhr){
                xhr.setRequestHeader('Authorization', readCookie("token"));
            },
            success: function(data, textStatus, jqXHR) {
                viewSuccess("Value: "+data);
            },
            error: function(jqXHR, textStatus, errorThrown) {
                viewError(jqXHR.responseText);
            }
        });
        e.preventDefault(e);
    });

    $('#delete_data').click (function(e) {
        var data = getFormData($('#data-form'));
        var err = "";
        clearError();
        clearSuccess();

        if (data.key == "") {
            viewError("Key is required field");
            return false;
        }

        key = data.key
        $.ajax({
            url: pref+'/value/'+data.key,
            method: 'DELETE',
            contentType: "application/json; charset=utf-8",
            beforeSend: function setAuthHeader(xhr){
                xhr.setRequestHeader('Authorization', readCookie("token"));
            },
            success: function(data, textStatus, jqXHR) {
                viewSuccess("Delete with key "+key);
            },
            error: function(jqXHR, textStatus, errorThrown) {
                viewError(jqXHR.responseText);
            }
        });
        e.preventDefault(e);
    });
});

function getHistory() {
    $.ajax({
        url: pref+'/history',
        method: 'GET',
        contentType: "application/json; charset=utf-8",
        dataType: 'json',
        beforeSend: function setAuthHeader(xhr){
            xhr.setRequestHeader('Authorization', readCookie("token"));
        },
        success: function(data, textStatus, jqXHR) {
            var content = '';
            $.each(data, function( index, value ) {
                var date = new Date(Date.parse(value.CreatedAt)).toUTCString();
                content += '<tr>';
                content += '<td>'+value.User.Login+'</td>';
                content += '<td>'+date+'</td>';
                content += '<td>'+value.Query+'</td>';
                content += '<td>'+value.Params+'</td>';
                content += '</tr>';
            });
            $('#history_content').html(content);
        },
        error: function(jqXHR, textStatus, errorThrown) {
            console.log(jqXHR.responseText);
        }
    });
}

function viewError(text) {
    $('#form-error').show().html(text);
}

function viewSuccess(text) {
    $('#form-success').show().html(text);
}

function clearError() {
    $('#form-error').hide().html('');
}

function clearSuccess() {
    $('#form-success').hide().html('');
}
