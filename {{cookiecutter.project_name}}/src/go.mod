module github.com/{{cookiecutter.github_org}}/{{cookiecutter.project_slug}}

require (
    github.com/rs/zerolog {{ cookiecutter.zerolog_version }}
	github.com/spf13/cobra {{ cookiecutter.cobra_version }}
	{% if cookiecutter.use_viper == "y" -%}github.com/spf13/viper {{ cookiecutter.viper_version }}{%- endif %}
)