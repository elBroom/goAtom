$(function() {
    $('#login-form-link').click(function(e) {
        $("#login-form").delay(100).fadeIn(100);
        $("#register-form").fadeOut(100);
        $('#register-form-link').removeClass('active');
        $(this).addClass('active');
        e.preventDefault();
    });
    $('#register-form-link').click(function(e) {
        $("#register-form").delay(100).fadeIn(100);
        $("#login-form").fadeOut(100);
        $('#login-form-link').removeClass('active');
        $(this).addClass('active');
        e.preventDefault();
    });

    $('#login-form').submit(function(e) {
        var data = getFormData($(this));
        var err = "";
        clearError("register");

        if (data.login == "") {
            err = "Login is required field";
        } else if (data.password == "") {
            err = "Password is required field";
        }

        if (err) {
            viewError("login", err);
            return false;
        }

        $.ajax({
            url: pref+'/login',
            method: 'POST',
            contentType: "application/json; charset=utf-8",
            dataType: 'json',
            data: JSON.stringify(data),
            success: function(data, textStatus, jqXHR) {
                debugger;
            },
            error: function(jqXHR, textStatus, errorThrown) {
                if (jqXHR.responseText == "ok") {
                    createCookie("token", jqXHR.getResponseHeader('Authorization'), 1);
                    window.location = index_link;
                } else {
                    viewError("login", jqXHR.responseText);
                }
            }
        });
        e.preventDefault(e);
    });

    $('#register-form').submit(function(e) {
        var data = getFormData($(this));
        var err = "";
        clearError("register");

        if (data.login == "") {
            err = "Login is required field";
        } else if (data.password == "") {
            err = "Password is required field";
        } else if (data.confirm_password == "") {
            err = "Confirm Password is required field";
        } else if (data.password != data.confirm_password ) {
            err = "Password and Confirm Password not equals";
        }

        if (err) {
            viewError("register", err);
            return false;
        }

        $.ajax({
            url: pref+'/user',
            method: 'POST',
            contentType: "application/json; charset=utf-8",
            dataType: 'json',
            data: JSON.stringify(data),
            success: function(data, textStatus, jqXHR) {
                debugger;
            },
            error: function(jqXHR, textStatus, errorThrown) {
                if (jqXHR.responseText == "ok") {
                    $('#login-form-link').trigger( "click" );
                } else {
                    viewError("register", jqXHR.responseText);
                }
            }
        });
        e.preventDefault(e);
    });
});

function viewError(context, text) {
    $('#'+context+'-form-error').show().html(text);
}

function clearError(context) {
    $('#'+context+'-form-error').hide().html('');
}
