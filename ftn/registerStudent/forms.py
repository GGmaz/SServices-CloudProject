from django import forms

from .models import Student

# Create your forms here.

class RegistrationForm(forms.Form):
    first_name = forms.CharField(label='Your first name')
    last_name = forms.CharField(label='Your last name')
    jmbg = forms.CharField(label='Your jmbg')
    index = forms.CharField(label='Your index')

    # write me method for save data in database
    def save(self):
        data = self.cleaned_data
        student = Student(first_name=data['first_name'], last_name=data['last_name'], jmbg=data['jmbg'], index=data['index'])
        student.save()

    def all(self):
        return Student.objects.all()
