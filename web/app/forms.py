from django import forms


class RegistrationForm(forms.Form):
    first_name = forms.CharField(max_length=100)
    last_name = forms.CharField(max_length=100)
    email = forms.EmailField(max_length=100)
    email_confirm = forms.EmailField(max_length=100)
    twitter_handle = forms.CharField(max_length=100, required=False)
    company = forms.CharField(max_length=100, required=False)
    position = forms.CharField(max_length=100, required=False)
    phone = forms.CharField(max_length=100, required=False)

    def clean(self):
        super(RegistrationForm, self).clean()
        email = self.cleaned_data.get('email')
        email_confirm = self.cleaned_data.get('email_confirm')
        if email != email_confirm:
            raise forms.ValidationError("Email confirmation does not match")

