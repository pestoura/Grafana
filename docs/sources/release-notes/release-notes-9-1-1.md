---
aliases: []
hide_menu: true
title: Release notes for Grafana 9.1.1
---

<!-- Auto generated by update changelog github action -->

# Release notes for Grafana 9.1.1

### Features and enhancements

- **Cloud Monitoring:** Support SLO burn rate. [#53710](https://github.com/grafana/grafana/pull/53710), [@itkq](https://github.com/itkq)
- **Schema:** Restore "hidden" in LegendDisplayMode. [#53925](https://github.com/grafana/grafana/pull/53925), [@academo](https://github.com/academo)
- **Timeseries:** Revert the timezone(s) property name change back to singular. [#53926](https://github.com/grafana/grafana/pull/53926), [@academo](https://github.com/academo)

### Bug fixes

- **Alerting:** Fix links in Microsoft Teams notifications. [#54003](https://github.com/grafana/grafana/pull/54003), [@grobinson-grafana](https://github.com/grobinson-grafana)
- **Alerting:** Fix notifications for Microsoft Teams. [#53810](https://github.com/grafana/grafana/pull/53810), [@grobinson-grafana](https://github.com/grobinson-grafana)
- **Alerting:** Fix width of Adaptive Cards in Teams notifications. [#53996](https://github.com/grafana/grafana/pull/53996), [@grobinson-grafana](https://github.com/grobinson-grafana)
- **ColorPickerInput:** Fix popover in disabled state. [#54000](https://github.com/grafana/grafana/pull/54000), [@Clarity-89](https://github.com/Clarity-89)
- **Decimals:** Fixes auto decimals to behave the same for positive and negative values. [#53960](https://github.com/grafana/grafana/pull/53960), [@JoaoSilvaGrafana](https://github.com/JoaoSilvaGrafana)
- **Loki:** Fix unique log row id generation. [#53932](https://github.com/grafana/grafana/pull/53932), [@gabor](https://github.com/gabor)
- **Plugins:** Fix file extension in development authentication guide. [#53838](https://github.com/grafana/grafana/pull/53838), [@pbzona](https://github.com/pbzona)
- **TimeSeries:** Fix jumping legend issue. [#53671](https://github.com/grafana/grafana/pull/53671), [@zoltanbedi](https://github.com/zoltanbedi)
- **TimeSeries:** Fix memory leak on viz re-init caused by KeyboardPlugin. [#53872](https://github.com/grafana/grafana/pull/53872), [@leeoniya](https://github.com/leeoniya)

### Plugin development fixes & changes

- **TimePicker:** Fixes relative timerange of less than a day not displaying. [#53975](https://github.com/grafana/grafana/pull/53975), [@JoaoSilvaGrafana](https://github.com/JoaoSilvaGrafana)
- **GrafanaUI:** Fixes ClipboardButton to always keep multi line content. [#53903](https://github.com/grafana/grafana/pull/53903), [@JoaoSilvaGrafana](https://github.com/JoaoSilvaGrafana)